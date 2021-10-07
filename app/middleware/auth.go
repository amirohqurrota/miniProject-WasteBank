package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		//ErrorHandlerWithContext: errors.New("forbidden access"),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(roleID int, role string) string {
	claims := JwtCustomClaims{
		roleID,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))
	return token
}

//parsing token string
func ParsingToken(tokenString, role string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret!@#$%"), nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	return claims, nil
}

func RoleValidation(tokenString, role string) bool {
	result, err := ParsingToken(tokenString, role)
	if err != nil {
		return false
	}
	roleLogin := result["role"]
	return roleLogin == role
}

//cel

// func GetUser(c echo.Context, role string) *JwtCustomClaims {
// 	roleLogin := c.Get(role).(*jwt.Token)
// 	claims := roleLogin.Claims.(*JwtCustomClaims)
// 	return claims
// }

// func RoleValidationCheck(roleLogin string) echo.MiddlewareFunc {
// 	return func(hf echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			claims := GetUser(c, roleLogin)

// 			if claims.Role == roleLogin {
// 				return hf(c)
// 			} else {
// 				return echo.ErrForbidden
// 			}
// 		}
// 	}
// }
