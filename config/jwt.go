package config

import "time"

type JwtConfig struct {
	Secret string
	Exp    time.Duration // 过期时间
	Alg    string        // 算法

}

func GetJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret: "jxc_jwt_session",
		Exp:    time.Hour * 24 * 3,
		Alg:    "HS256",
	}
}
