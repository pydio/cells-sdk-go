package transport

type StaticTokenProvider struct {
	staticToken string
}

func NewStaticTokenProvider(token string) TokenProvider {
	return &StaticTokenProvider{staticToken: token}
}

func (t *StaticTokenProvider) Retrieve() (string, error) {
	return t.staticToken, nil
}

func (t *StaticTokenProvider) Expired() bool {
	return false
}
