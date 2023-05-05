package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"mall/consts"
)

var jwtSecret = []byte("FanOne")

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(id uint, username string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.AccessTokenExpireDuration)
	rtExpireTime := nowTime.Add(consts.RefreshTokenExpireDuration)
	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}
	// 加密并获得完整的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "mall",
	}).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ParseRefreshToken 验证用户token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	_, err = jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return "", err
	})

	var claims Claims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return "", err
	})

	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenerateToken(claims.ID, claims.Username)
	}

	return
}

// EmailClaims
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken 签发邮箱验证Token
func GenerateEmailToken(userID, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * time.Minute)
	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cmall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken 验证邮箱验证token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
