package auth

import (
	"context"
	"fmt"
	"github.com/go-frame/config"
	"github.com/go-frame/internals/async_worker"
	"github.com/go-frame/internals/entity"
	"github.com/go-frame/internals/entity/apperror"
	"github.com/go-frame/internals/entity/httpentity"
	"github.com/go-frame/internals/entity/proto"
	"github.com/go-frame/internals/lib"
	"github.com/go-frame/internals/lib/logger"
	"github.com/golang-jwt/jwt"
	"github.com/mssola/useragent"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strings"
	"time"
)

type Service struct {
	repo             Repository
	logger           logger.Logger
	config           *config.Config
	usersServiceRepo UsersServiceRepo
	redisClient      *redis.Client
	workerPool       *async_worker.WorkerPool
	workerService    WorkerService
}

func NewService(repository Repository, logger logger.Logger, config *config.Config, usersServiceRepo UsersServiceRepo, redisClient *redis.Client) *Service {
	workerPool := async_worker.NewWorkerPool(100)
	if workerPool == nil {
		logger.Error("Failed to initialize loginAttemptWorker")
		return nil
	}

	service := &Service{
		repo:             repository,
		logger:           logger,
		config:           config,
		usersServiceRepo: usersServiceRepo,
		redisClient:      redisClient,
		workerPool:       workerPool,
		workerService: WorkerService{
			loginAttemptWorker: workerPool,
			logger:             repository,
		},
	}

	// Start the login attempt worker
	ctx := context.Background()
	service.workerPool.Start(ctx, 10)

	return service
}

// LoginAttemptTask is a task for logging login attempts.
type LoginAttemptTask struct {
	async_worker.TaskStatus
	Attempt entity.LoginAttempt
	Logger  LoginAttemptLogger
}

// UpdateLoginAttemptTask is another example task that implements the Task interface.
type UpdateLoginAttemptTask struct {
	async_worker.TaskStatus
	AttemptID  int
	UpdateFunc func(ctx context.Context, attemptID int) error
}

// Process processes the login attempt task.
func (t *LoginAttemptTask) Process(ctx context.Context) error {
	// Execute common pre-processing steps
	if err := t.PreProcess(ctx); err != nil {
		return fmt.Errorf("pre-processing failed: %v", err)
	}
	defer func() {
		// Execute common post-processing steps (deferred to ensure it runs even on error)
		if err := t.PostProcess(ctx); err != nil {
			fmt.Printf("[%s] post-processing failed: ", err)
		}
	}()

	// Perform the specific task logic
	if err := t.Logger.LogLoginAttempt(ctx, &t.Attempt); err != nil {
		t.SetErrorStatus(err)
		return err
	}
	return nil
}

func (t *UpdateLoginAttemptTask) Process(ctx context.Context) error {
	// Execute common pre-processing steps
	if err := t.PreProcess(ctx); err != nil {
		return fmt.Errorf("pre-processing failed: %v", err)
	}
	defer func() {
		// Execute common post-processing steps (deferred to ensure it runs even on error)
		if err := t.PostProcess(ctx); err != nil {
			fmt.Printf("[%s] post-processing failed: %v\n", err)
		}
	}()

	// Perform the specific task logic
	if err := t.UpdateFunc(ctx, t.AttemptID); err != nil {
		t.SetErrorStatus(err)
		return err
	}
	return nil
}

// Service represents the main service.
type WorkerService struct {
	loginAttemptWorker *async_worker.WorkerPool
	logger             LoginAttemptLogger
}

const (
	maxAttempts     = 5
	lockoutDuration = 15 * time.Minute
)

func (s *Service) trackLoginAttempts(ctx context.Context, email string) (bool, time.Duration, error) {
	redisKey := fmt.Sprintf("login_attempts:%s", email)

	// Use Redis to track attempts
	attempts, err := s.redisClient.Get(ctx, redisKey).Int()
	if err != nil && err != redis.Nil {
		return false, 0, err
	}

	// If max attempts reached, return lockout status and remaining time
	if attempts >= maxAttempts {
		ttl, err := s.redisClient.TTL(ctx, redisKey).Result()
		if err != nil {
			return false, 0, err
		}
		return true, ttl, nil
	}

	// Increment attempt count in Redis
	s.redisClient.Incr(ctx, redisKey)
	s.redisClient.Expire(ctx, redisKey, lockoutDuration)

	return false, 0, nil
}

func (s *Service) extractDeviceInfo(userAgent string) (device, os, osVersion string) {
	ua := useragent.New(userAgent)
	name, version := ua.Browser()

	osInfo := ua.OS()
	if ua.Mobile() {
		device = "Mobile"
	} else if ua.Bot() {
		device = "Bot"
	} else {
		device = name
	}

	osParts := strings.Split(osInfo, " ")
	if len(osParts) > 1 {
		os = osParts[0]
		osVersion = strings.Join(osParts[1:], " ")
	} else {
		os = osInfo
		osVersion = version
	}

	return device, os, osVersion
}

func (s *Service) insertLoginAttempt(email, ip, userAgent string, status bool, message string) {
	device, os, osVersion := s.extractDeviceInfo(userAgent)
	taskStatus := async_worker.NewTaskStatus()
	task := &LoginAttemptTask{
		TaskStatus: *taskStatus,
		Attempt: entity.LoginAttempt{
			Email:       email,
			IPAddress:   ip,
			AttemptTime: time.Now(),
			UserAgent:   userAgent,
			Device:      device,
			Os:          os,
			OsVersion:   osVersion,
			Successful:  status,
			Reason:      message,
		},
		Logger: s.workerService.logger,
	}

	go func() {
		if err := s.workerPool.Enqueue(task); err != nil {
			s.logger.Error("Failed to enqueue login attempt for logging:", err)
		}
	}()
}

func (s *Service) updateLoginAttempt(attemptID int) {
	taskStatus := async_worker.NewTaskStatus()
	// Update login attempt
	updateTask := &UpdateLoginAttemptTask{
		TaskStatus: *taskStatus,
		AttemptID:  attemptID,
		UpdateFunc: s.repo.UpdateLoginAttempt,
	}

	go func() {
		if err := s.workerPool.Enqueue(updateTask); err != nil {
			s.logger.Error("Failed to enqueue login attempt update:", err)
		}
	}()
}

func (s *Service) resetLoginAttempts(ctx context.Context, email string) error {
	redisKey := fmt.Sprintf("login_attempts:%s", email)
	return s.redisClient.Del(ctx, redisKey).Err()
}
func setCookie(w http.ResponseWriter, name, value string, duration time.Duration) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(duration),
	})
}

func (s *Service) getClaimDataByUserID(ctx context.Context, userID string) (*proto.UserClaimData, error) {
	in := &proto.ClaimDataByUserIDRequest{
		UserId: userID,
	}
	res, err := s.usersServiceRepo.GetClaimDataByUserID(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) generateTokenForUser(claimData *proto.UserClaimData) (httpentity.UserTokenResponse, error) {
	claims := entity.JwtClaim{
		CustomClaim: s.generateCustomClaim(claimData),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * config.ACCESSTOKEN_VALIDITY_MINS).Unix(),
		},
	}
	token := claims.NewToken()
	refreshToken, _ := claims.RefreshToken()

	res := httpentity.UserTokenResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresIn:    config.ACCESSTOKEN_VALIDITY_MINS,
	}
	return res, nil
}

func (s *Service) generateCustomClaim(claimData *proto.UserClaimData) entity.CustomClaim {
	userGroups := make([]entity.UserGroups, len(claimData.UserGroups))
	for i, group := range claimData.UserGroups {
		userGroups[i] = entity.UserGroups{
			UserGroupId: group.UserGroupId,
		}
	}

	return entity.CustomClaim{
		UserId:         claimData.UserId,
		FirstName:      claimData.FirstName,
		LastName:       claimData.LastName,
		LocationId:     int(claimData.LocationId),
		LocationTypeId: int(claimData.LocationTypeId),
		Email:          claimData.Email,
		UserGroups:     userGroups,
	}
}

func (s *Service) handleTokenGenerationAndCookies(ctx context.Context, claimData *proto.UserClaimData, w http.ResponseWriter) error {
	tokenGenerateRes, err := s.generateTokenForUser(claimData)
	if err != nil {
		return err
	}
	redisKey := fmt.Sprintf("refresh_token:%s", claimData.UserId)
	err = s.redisClient.Set(ctx, redisKey, tokenGenerateRes.RefreshToken, time.Minute*time.Duration(config.REFRESHTOKEN_VALIDITY_MINS)).Err()
	if err != nil {
		return err
	}
	passphrase := s.config.PassPhrase
	encryptedAccessToken, err := lib.Encrypt(tokenGenerateRes.AccessToken, passphrase)
	if err != nil {
		return err
	}
	setCookie(w, "access_token", encryptedAccessToken, time.Minute*time.Duration(tokenGenerateRes.ExpiresIn))
	setCookie(w, "refresh_token", tokenGenerateRes.RefreshToken, time.Minute*time.Duration(config.REFRESHTOKEN_VALIDITY_MINS))
	return nil
}

func (s *Service) EmailLogin(ctx context.Context, payloads *httpentity.EmailLoginRequest, w http.ResponseWriter, ip string, userAgent string) error {

	lockout, ttl, err := s.trackLoginAttempts(ctx, payloads.Email)
	if err != nil {
		return apperror.New(http.StatusInternalServerError, "internal_server_error", "An error occurred while processing your request")
	}
	if lockout {
		s.insertLoginAttempt(payloads.Email, ip, userAgent, false, "TOO_MANY_ATTEMPTS")
		return apperror.New(http.StatusTooManyRequests, "too_many_attempts", fmt.Sprintf("Too many login attempts. Please try again in %v.", ttl))
	}

	requestBody := &proto.VerifyPasswordRequest{
		Email:    payloads.Email,
		Password: payloads.Password,
	}
	res, err := s.usersServiceRepo.VerifyPasswordAndGetClaimData(ctx, requestBody)

	if err != nil {
		fmt.Println("error.....", err)
		return err
	}

	if !res.Match {
		s.insertLoginAttempt(payloads.Email, ip, userAgent, false, "PASS_OR_EMAIL_DIDNT_MATCH")
		s.updateLoginAttempt(31)
		return apperror.New(http.StatusUnauthorized, "pass_or_email_didnt_match", "Your Password or email did not match")
	}

	if err := s.resetLoginAttempts(ctx, payloads.Email); err != nil {
		s.logger.Error("Failed to reset login attempts: ", err)
	}
	s.insertLoginAttempt(payloads.Email, ip, userAgent, true, "LOGIN_ATTEMPTS_SUCCESSFULLY")
	return s.handleTokenGenerationAndCookies(ctx, res.ClaimData, w)

}

func (s *Service) RefreshToken(ctx context.Context, payloads *httpentity.UserRefreshRequest, w http.ResponseWriter) error {
	redisKey := fmt.Sprintf("refresh_token:%s", payloads.UserId)
	storedRefreshToken, err := s.redisClient.Get(ctx, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			return apperror.New(http.StatusUnauthorized, "invalid_refresh_token", "Invalid refresh token")
		}
		return err
	}

	if storedRefreshToken != payloads.RefreshToken {
		return apperror.New(http.StatusUnauthorized, "invalid_refresh_token", "Invalid refresh token")
	}

	// Assuming you have a function to get claim data by user ID
	claimData, err := s.getClaimDataByUserID(ctx, payloads.UserId)
	if err != nil {
		return err
	}

	return s.handleTokenGenerationAndCookies(ctx, claimData, w)
}

func (s *Service) Close() {
	s.workerPool.Stop()
}
