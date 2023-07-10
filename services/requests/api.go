package requests

import (
	"context"
	"time"
)

func (service *Service) ThreadSafe() *Service {
	service.threadSafe = true
	return service
}

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

func (service *Service) Retries(retries int) *Service {
	service.retries = retries
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

		ctx, cancel = service.Ctx()
		defer cancel()
	}

	return service.SendCtx(ctx)
}

func (service *Service) SendCtx(ctx context.Context) ([]byte, error) {
	response, err := service.sendRequest(ctx)
	if err != nil || (isError(service.response.StatusCode) && service.retries > 0) {
		for i := 0; i < service.retries; i++ {
			response, err = service.sendRequest(ctx)
			if err == nil && isOk(service.response.StatusCode) {
				return response, nil
			}
		}

		return nil, err
	}

	return response, nil
}

func (service *Service) SendStatus() ([]byte, int, error) {
	response, err := service.Send()

	return response, service.response.StatusCode, err
}

func (service *Service) SendStatusCtx(ctx context.Context) ([]byte, int, error) {
	response, err := service.SendCtx(ctx)

	return response, service.response.StatusCode, err
}
