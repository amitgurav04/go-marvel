package marvel

import (
	"strconv"
	"time"
	"crypto/md5"
	"encoding/hex"
)

// AuthParameters: Query parameter values provided for authentication
type AuthParameters struct {
	ApiKey string `url:"apikey"`
	Ts     string `url:"ts,omitempty"`
	Hash   string `url:"hash,omitempty"`
}

// Authenticator: Interface for providing AuthParameters
type Authenticator interface {
	GetAuth() *AuthParameters
}

// Auth: Contains API keys and timestamp for server side authentication
type Auth struct {
	publicKey  string
	privateKey string
	ts         string
}

// NewAuth returns a Auth with currentTimeStamp
func NewAuth(pubKey, priKey string) *Auth {
	currentTimeStamp := strconv.FormatInt(time.Now().UnixNano(), 16)

	return &Auth{
		publicKey:  pubKey,
		privateKey: priKey,
		ts:         currentTimeStamp,
	}
}

// GetAuth implements Authenticator interface
func (auth *Auth) GetAuth() *AuthParameters {
	hr := md5.New()
	hr.Write([]byte(auth.ts + auth.privateKey + auth.publicKey))
	hash := hex.EncodeToString(hr.Sum(nil))

	return &AuthParameters{
		ApiKey: auth.publicKey,
		Ts:     auth.ts,
		Hash:   hash,
	}
}
