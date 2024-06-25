// version 1 package for handler
package v1

import (
	"github.com/go-frame/internals/entity"
	"github.com/go-frame/internals/entity/httpentity"
	"github.com/go-frame/internals/lib/logger"
	"github.com/go-frame/internals/service/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	Service *auth.Service
	Logger  logger.Logger
}

func NewAuthHandler(authService *auth.Service, logger logger.Logger) *AuthHandler {
	return &AuthHandler{
		Service: authService,
		Logger:  logger,
	}
}

func (h *AuthHandler) MapAuthRoutes(authGroup *echo.Group) {
	authGroup.POST("/login", h.LoginHandler)
	authGroup.POST("/token/refresh", h.RefreshToken)
}

func (h *AuthHandler) LoginHandler(c echo.Context) error {
	payLoads := httpentity.EmailLoginRequest{}
	err := c.Bind(&payLoads)
	if err != nil {
		return handleInvalidDataError(c, err)
	}
	validationErrors := payLoads.Validate()
	if validationErrors != nil {
		return handleValidationError(c, validationErrors)
	}

	err = h.Service.EmailLogin(c.Request().Context(), &payLoads, c.Response().Writer, c.RealIP(), c.Request().UserAgent())
	if err != nil {
		return c.JSON(http.StatusOK, &httpentity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &httpentity.Response{
		Success: true,
		Message: "Login Successfully",
	})
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()
	data := httpentity.UserRefreshRequest{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Invalid data",
		})
	}
	validationErrors := data.Validate()
	if validationErrors != nil {
		return handleValidationError(c, validationErrors)
	}
	err = h.Service.RefreshToken(ctx, &data, c.Response().Writer)
	if err != nil {
		return c.JSON(http.StatusOK, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "Token refresh successfully",
	})
}
