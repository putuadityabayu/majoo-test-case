package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	JWT_ISSUER                = "https://majoo.id"
	APPLICATION_NAME          = "Majoo Test Cases 1"
	JWT_SIGNATURE_KEY         = []byte("the secret of kalimdor")
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
)
