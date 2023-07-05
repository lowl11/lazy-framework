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

func (service *Service) XML() *Service {
	service.isXml = true

	return service
}

func (service *Service) JSON() *Service {
	service.isXml = false

	return service
}

func (service *Service) SslTrust() *Service {
	service.sslCheck = true

	return service
}

func (service *Service) NoProxy() *Service {
	service.noProxy = true

	return service
}

func (service *Service) Timeout(seconds time.Duration) *Service {
	service.timeout = seconds
	service.customTimeout = true

	return service
}

func (service *Service) Header(key, value string) *Service {
	if _, ok := service.headers[key]; ok {
		service.headers[key] = append(service.headers[key], value)
	} else {
		service.headers[key] = []string{value}
	}

	return service
}

func (service *Service) Headers(headers map[string][]string) *Service {
	service.headers = headers

	return service
}

func (service *Service) Cookie(key, value string) *Service {
	service.cookies[key] = value

	return service
}

func (service *Service) Cookies(cookies map[string]string) *Service {
	service.cookies = cookies

	return service
}

func (service *Service) BasicAuth(username, password string) *Service {
	service.username = username
	service.password = password
	service.isBasicAuth = true

	return service
}

func (service *Service) Send() ([]byte, error) {
	ctx := context.Background()
	if service.customTimeout {
		var cancel func()

		ctx, cancel = service.Ctx(service.timeout)
		defer cancel()
	}

	return service.SendCtx(ctx)
}

func (service *Service) SendCtx(ctx context.Context) ([]byte, error) {
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

	if service.body != nil {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, bodyBuffer)
	} else {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, nil)
	}
	if err != nil {
		return nil, err
	}

	service.request = request

	service.fillHeaders()
	service.fillCookies()
	service.setBasicAuth()

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !service.sslCheck,
			},
		},
	}

	if service.noProxy {
		client.Transport.(*http.Transport).Proxy = nil
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Error(err, "Close response body error")
		}
	}()

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	service.status = response.StatusCode

	return responseInBytes, nil
}

func (service *Service) SendStatus() ([]byte, int, error) {
	response, err := service.Send()

	return response, service.status, err
}

func (service *Service) SendStatusCtx(ctx context.Context) ([]byte, int, error) {
	response, err := service.SendCtx(ctx)

	return response, service.status, err
}
