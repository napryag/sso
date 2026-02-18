package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/napryag/sso/internal/domain/models"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	stringToken, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return stringToken, nil
}
