package service_token

import (
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/philips-software/go-hsdp-api/iam"
	"net/http"
)

func init() {
	caddy.RegisterModule(ServiceToken{})
	httpcaddyfile.RegisterHandlerDirective("service_token", parseCaddyfile)
}

type ServiceToken struct {
	Region      string `json:"region,omitempty"`
	Environment string `json:"environment,omitempty"`
	ServiceId   string `json:"service_id,omitempty"`
	ServiceKey  string `json:"service_key,omitempty"`
}

func (ServiceToken) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.service_token",
		New: func() caddy.Module { return new(ServiceToken) },
	}
}

func (m *ServiceToken) Provision(ctx caddy.Context) error {
	return nil
}

func (m *ServiceToken) Validate() error {
	return nil
}

func (m ServiceToken) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	iamClient, err := iam.NewClient(http.DefaultClient, &iam.Config{
		Region:      m.Region,
		Environment: m.Environment,
	})
	if err != nil {
		println("error initializing IAM client", "error", err)
		return err
	}
	println("logging in", "serviceID", m.ServiceId)
	err = iamClient.ServiceLogin(iam.Service{
		ServiceID:  m.ServiceId,
		PrivateKey: m.ServiceKey,
	})
	if err != nil {
		println("error logging in", "error", err)
		return err
	}

	token, _ := iamClient.Token()

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return next.ServeHTTP(w, r)
}

func (m *ServiceToken) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for d.NextBlock(0) {
			if err := m.setField(d); err != nil {
				return err
			}
		}
	}
	return nil
}

// setField maps dispenser tokens to struct fields and assigns values accordingly.
func (m *ServiceToken) setField(d *caddyfile.Dispenser) error {
	if !d.NextArg() {
		return d.ArgErr()
	}

	value := d.Val()
	switch d.Val() {
	case "region":
		m.Region = value
	case "environment":
		m.Environment = value
	case "service_id":
		m.ServiceId = value
	case "service_key":
		m.ServiceKey = value
	default:
		return d.Errf("unexpected token '%s' in service_token block", d.Val())
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
