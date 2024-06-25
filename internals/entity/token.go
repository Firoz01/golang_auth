package entity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserGroups struct {
	UserGroupId string `json:"user_group_id"`
}

type CustomClaim struct {
	UserId         string       `json:"user_id"`
	FirstName      string       `json:"first_name"`
	LastName       string       `json:"last_name"`
	LocationId     int          `json:"location_id"`
	LocationTypeId int          `json:"location_type_id"`
	Email          string       `json:"email"`
	UserGroups     []UserGroups `json:"user_groups"`
}
type JwtClaim struct {
	CustomClaim
	jwt.StandardClaims
}

func (claims *JwtClaim) NewToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte("secret"))
	return t
}
func (claims *JwtClaim) RefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
