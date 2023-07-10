package requests

import (
	"net/http"
	"time"
)

type Service struct {
	// basic info
	method string
	url    string
	body   any

	// request container
	headers map[string][]string
	cookies map[string]string

	// body format
	isXml bool

	// timeout
	timeout       time.Duration
	customTimeout bool

	// retries
	retries int

	// auth
	isBasicAuth bool
	username    string
	password    string

	// network
	sslCheck bool
	noProxy  bool

	// req/resp
	request  *http.Request
	response *http.Response

	responseBody []byte
}

func New(method, url string, body any) *Service {
	return &Service{
		method: method,
		url:    url,
		body:   body,

		headers: make(map[string][]string),
		cookies: make(map[string]string),

		responseBody: []byte{},
	}
}

func NewSoap(url string, body any) *Service {
	return &Service{
		method: http.MethodPost,
		url:    url,
		body:   body,
		isXml:  true,

		headers: make(map[string][]string),
		cookies: make(map[string]string),
	}
}
