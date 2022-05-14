//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package api

type RequestHandler interface {
	Get(url string, headers map[string]string) (Response, error)
}

type Response interface {
	BodyBytes() ([]byte, error)
	IsSuccess() bool
}
