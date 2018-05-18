package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type TokenStore struct {
	internalCache *cache.Cache
}

func NewTokenStore() *TokenStore {
	t := &TokenStore{
		internalCache: cache.New(20*time.Minute, 10*time.Minute),
	}
	return t
}

func (t *TokenStore) Store(c *SdkConfig, token string, expiry time.Duration) {
	//fmt.Println("[Auth] Storing token with expiration ", expiry)
	t.internalCache.Set(t.computeKey(c), token, expiry)
}

func (t *TokenStore) TokenFor(c *SdkConfig) string {
	if token, ok := t.internalCache.Get(t.computeKey(c)); ok {
		return token.(string)
	} else {
		return ""
	}
}

func (t *TokenStore) computeKey(c *SdkConfig) string {
	s := fmt.Sprintf("%s-%s-%s-%s-%s", c.Url, c.ClientKey, c.ClientSecret, c.User, c.Password)
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
