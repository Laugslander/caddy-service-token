# caddy-service-token

Plugin for [Caddy](https://caddyserver.com/) that automatically injects an [HSP](https://www.hsdp.io/) IAM service
identity access token as request header.  
Useful when running Caddy as a reverse proxy for a HSP service in a microservice architecture.

## Configuration

```caddyfile
{
    order service_token first
}

:8080 {
    service_token {
        region "<region>"
        environment "<environment>"
        service_id "<service_id>"
        service_key "<service_key>"
    }

    reverse_proxy <hsp_service> {
        header_up Host {upstream_hostport}
    }
}
```

Populate all variables marked with `<variable>`.

## Development

Install [xcaddy](https://github.com/caddyserver/xcaddy):

```shell
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

Run the plugin locally:

```shell
xcaddy run
```

Build caddy with this plugin:

```shell
xcaddy build --with github.com/Laugslander/service-caddy-token
```

For more information, see [Extending Caddy](https://caddyserver.com/docs/extending-caddy).

## License

License is Apache 2.0