FROM alpine:3.18

COPY ./caddy /usr/bin/caddy

RUN chmod +x /usr/bin/caddy

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]