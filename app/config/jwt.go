package config

var (
	jwtSecret string
)

func InitJWTSecret(secret string) {
	jwtSecret = secret
}

func GetJWTSecret() string {
	return jwtSecret
}
