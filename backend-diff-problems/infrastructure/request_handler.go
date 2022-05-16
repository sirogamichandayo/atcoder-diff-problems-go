package infrastructure

import (
	"bytes"
	"compress/gzip"
	"diff-problems/interfaces/api"
	"net/http"
)

type RequestHandler struct {
}

func NewRequestHandler() api.RequestHandler {
	clientHandler := new(RequestHandler)

	return clientHandler
}

func (handler *RequestHandler) Get(url string, headers map[string]string) (api.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Set("Accept-Encoding", "gzip")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}

	output := bytes.Buffer{}
	if _, err := output.ReadFrom(reader); err != nil {
		return nil, err
	}
	return Response{res.StatusCode, output.Bytes()}, err
}

type Response struct {
	statusCode int
	bodyBytes  []byte
}

func (res Response) IsSuccess() bool {
	return 200 <= res.statusCode && res.statusCode < 300
}

func (res Response) BodyBytes() []byte {
	return res.bodyBytes
}
