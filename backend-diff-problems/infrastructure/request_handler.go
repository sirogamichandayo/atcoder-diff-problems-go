package infrastructure

import (
	"bytes"
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
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	output := bytes.Buffer{}
	if _, err := output.ReadFrom(res.Body); err != nil {
		return nil, err
	}
	return Response{res.StatusCode, output.Bytes()}, err
}
