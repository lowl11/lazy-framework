package grpc_service

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"time"
)

func (service *Service) connect() (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	if service.creds != nil {
		creds = service.creds
	}

	options := service.opts
	options = append(options, grpc.WithTransportCredentials(creds))

	// set no proxy
	if service.noProxy {
		options = append(options, grpc.WithContextDialer(func(ctx context.Context, address string) (net.Conn, error) {
			// Create a http.Transport with the no_proxy setting
			transport := &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: !service.sslCheck,
				},
			}

			dialer := &net.Dialer{
				Timeout:   time.Second * 30,
				KeepAlive: time.Second * 30,
			}

			transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.DialContext(ctx, network, addr)
			}

			// Dial the address using the custom transport
			return transport.DialContext(ctx, "tcp", address)
		}))
	}

	ctx, cancel := context.WithTimeout(context.Background(), service.timeout)
	defer cancel()

	connection, err := grpc.DialContext(ctx, service.host, options...)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (service *Service) asd() {
	//
}
