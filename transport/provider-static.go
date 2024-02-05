package transport

//import (
//	"context"
//	"github.com/pydio/cells-sdk-go/v5"
//)
//
//type StaticTokenProvider struct {
//	staticToken string
//}
//
//func NewStaticTokenProvider(token string) cells_sdk.TokenProvider {
//	return &StaticTokenProvider{staticToken: token}
//}
//
//func (t *StaticTokenProvider) Retrieve(_ context.Context) (string, error) {
//	return t.staticToken, nil
//}
//
//func (t *StaticTokenProvider) Expired() bool {
//	return false
//}
