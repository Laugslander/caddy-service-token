package service_token

import (
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"net/http"
)

func init() {
	caddy.RegisterModule(ServiceToken{})
}

type ServiceToken struct{}

func (ServiceToken) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.service-token",
		New: func() caddy.Module { return new(ServiceToken) },
	}
}

func (h ServiceToken) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	fmt.Fprintf(w, "Hello from my plugin!")

	return next.ServeHTTP(w, r)
}

var (
	_ caddy.Module                = (*ServiceToken)(nil)
	_ caddyhttp.MiddlewareHandler = (*ServiceToken)(nil)
)
