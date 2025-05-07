package transport

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/pydio/cells-sdk-go/v4/apiv1"
)

var (
	pCache     map[string]apiv1.TokenProvider
	pCacheLock sync.Mutex
)

func initCache() {
	if pCache == nil {
		pCache = make(map[string]apiv1.TokenProvider)
	}
}

func WithProviderCache(provider apiv1.TokenProvider, c *apiv1.SdkConfig) apiv1.TokenProvider {
	initCache()
	pCacheLock.Lock()
	defer pCacheLock.Unlock()
	key := cacheKey(c)
	if cp, o := pCache[key]; o {
		return cp
	} else {
		pCache[key] = provider
		return provider
	}
}

func cacheKey(c *apiv1.SdkConfig) string {
	var s string
	if c.IdToken != "" {
		s = fmt.Sprintf("%s-%s", c.Url, c.User)
	} else {
		s = fmt.Sprintf("%s-%s-%s", c.Url, c.User, c.Password)
	}
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
