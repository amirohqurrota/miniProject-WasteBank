package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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

// GetUser from jwt ...
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("admin").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

// func ParsingToken(tokenString string) (jwt.MapClaims, error) {
// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("<YOUR VERIFICATION KEY>"), nil
// 	})
// 	if err != nil {
// 		return nil, errors.New("something wrong in your token request")
// 	}
// 	// decoded claims
// 	for key, val := range claims {
// 		fmt.Printf("Key: %v, value: %v\n", key, val)
// 	}
// 	fmt.Println("token : ", token.Header)
// 	return claims, nil
// }

func RoleValidation(role string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)
			if claims.Role == role {
				return hf(c)
			} else {
				return echo.ErrBadRequest
			}
		}
	}
}
