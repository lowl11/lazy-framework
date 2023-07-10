package requests

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/lazylog/log"
	"io/ioutil"
	"net/http"
	"time"
)

func (service *Service) sendRequest(ctx context.Context) ([]byte, error) {
	// parse body & convert to JSON/XML
	var bodyBuffer *bytes.Buffer
	if service.body != nil {
		if value, ok := service.body.([]byte); ok {
			bodyBuffer = bytes.NewBuffer(value)
		} else {
			var bodyInBytes []byte
			var err error
			if !service.isXml {
				bodyInBytes, err = json.Marshal(service.body)
			} else {
				bodyInBytes, err = xml.Marshal(service.body)
			}
			if err != nil {
				return nil, err
			}

			bodyBuffer = bytes.NewBuffer(bodyInBytes)
		}
	}

	var request *http.Request
	var err error

	// create request
	if service.body != nil {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, bodyBuffer)
	} else {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, nil)
	}
	if err != nil {
		return nil, err
	}

	// store request
	service.request = request

	// set request info
	service.fillHeaders()
	service.fillCookies()
	service.setBasicAuth()

	// create http client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !service.sslCheck,
			},
		},
	}

	// set no proxy
	if service.noProxy {
		client.Transport.(*http.Transport).Proxy = nil
	}

	// sending request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Error(err, "Close response body error")
		}
	}()

	// parse response
	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// store response
	service.response = response
	service.responseBody = responseInBytes

	return responseInBytes, nil
}

func (service *Service) fillHeaders() {
	if service.headers == nil {
		service.request.Header.Add("Connection", "keep-alive")
		return
	}

	for key, value := range service.headers {
		for _, item := range value {
			service.request.Header.Add(key, item)
		}
	}

	service.request.Header.Add("Connection", "keep-alive")
}

func (service *Service) fillCookies() {
	if service.cookies == nil {
		return
	}

	for key, value := range service.cookies {
		service.request.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}
}

func (service *Service) setBasicAuth() {
	if !service.isBasicAuth {
		return
	}

	service.request.SetBasicAuth(service.username, service.password)
}

func (service *Service) Ctx() (context.Context, func()) {
	defaultTimeout := time.Second * 10
	if service.customTimeout {
		defaultTimeout = service.timeout
	}

	return context.WithTimeout(context.Background(), defaultTimeout)
}

func isOk(code int) bool {
	return code >= http.StatusOK && code <= http.StatusIMUsed
}

func isError(code int) bool {
	return code >= http.StatusBadRequest
}
