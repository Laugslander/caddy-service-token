package service_token

import (
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"net/http"
)

func init() {
	caddy.RegisterModule(ServiceToken{})
	httpcaddyfile.RegisterHandlerDirective("service_token", parseCaddyfile)
}

type ServiceToken struct {
	ServiceId  string `json:"service_id,omitempty"`
	ServiceKey string `json:"service_key,omitempty"`
}

func (ServiceToken) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.service_token",
		New: func() caddy.Module { return new(ServiceToken) },
	}
}

func (m *ServiceToken) Provision(ctx caddy.Context) error {
	// TODO: Generate service token

	return nil
}

func (m *ServiceToken) Validate() error {
	// TODO: Validate inputs

	return nil
}

func (m ServiceToken) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	authorizationHeader := fmt.Sprintf("Bearer %s", "lala")

	r.Header.Set("Authorization", authorizationHeader)

	return next.ServeHTTP(w, r)
}

func (m *ServiceToken) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for d.NextBlock(0) {
			switch d.Val() {
			case "service_id":
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.ServiceId = d.Val()
			case "service_key":
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.ServiceKey = d.Val()
			default:
				return d.Errf("unexpected token '%s' in service_token block", d.Val())
			}
		}
	}
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m ServiceToken
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

var (
	_ caddy.Validator             = (*ServiceToken)(nil)
	_ caddyhttp.MiddlewareHandler = (*ServiceToken)(nil)
	_ caddyfile.Unmarshaler       = (*ServiceToken)(nil)
	_ caddy.Module                = (*ServiceToken)(nil)
)
