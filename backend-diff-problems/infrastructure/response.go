package infrastructure

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
