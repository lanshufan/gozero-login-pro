package pkg

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaim struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(userId, actTime int64) []byte {
	signKey := []byte("zeroAccessKey")
	// token生成、过期时间
	now := jwt.NewNumericDate(time.Now())
	exp := jwt.NewNumericDate(time.Unix(time.Now().Unix()+actTime, 0))
	// token信息
	mc := MyClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "lanshufan",
			Subject:   "zeroLogin",
			NotBefore: now,
			ExpiresAt: exp,
		},
	}
	// 签发token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	s, _ := token.SignedString(signKey)

	return []byte(s)
}

// ParseToken 分析token
func ParseToken(tokenStr []byte) (*jwt.Token, error) {
	signKey := []byte("zeroAccessKey")
	t, err := jwt.ParseWithClaims(string(tokenStr), &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	return t, err
}

// ValidToken 验证token
func ValidToken(tokenStr []byte) bool {
	t, err := ParseToken(tokenStr)
	if _, ok := t.Claims.(*MyClaim); !ok || !t.Valid || err != nil {
		return false
	}

	return true
}

// GetTokenId 返回Token id，无法验证则返回0
func GetTokenId(tokenStr []byte) int64 {
	t, err := ParseToken(tokenStr)
	if claim, ok := t.Claims.(*MyClaim); !ok || !t.Valid || err != nil {
		return 0
	} else {
		return claim.UserId
	}
}
