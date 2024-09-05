package jwt

import "os"

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtExpiration = os.Getenv("JWT_EXPIRATION")
