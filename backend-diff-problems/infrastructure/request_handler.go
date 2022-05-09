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

	req.Header.Set("Accept-Encoding", "gzip")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return Response{res}, err
}

type Response struct {
	Response *http.Response
}

func (res Response) BodyBytes() ([]byte, error) {
	reader, err := gzip.NewReader(res.Response.Body)
	if err != nil {
		return nil, err
	}

	output := bytes.Buffer{}
	if _, err := output.ReadFrom(reader); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
