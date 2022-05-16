package infrastructure

import (
	"bytes"
	"compress/gzip"
	"diff-problems/interfaces/api"
	"net/http"
)

type GzipRequestHandler struct {
}

func NewGzipRequestHandler() api.RequestHandler {
	clientHandler := new(GzipRequestHandler)

	return clientHandler
}

func (handler *GzipRequestHandler) Get(url string, headers map[string]string) (api.Response, error) {
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
