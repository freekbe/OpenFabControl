#!/bin/bash
set -e

# Default values
ENABLE_SSL_ON_NGINX=${ENABLE_SSL_ON_NGINX:-"true"}

if [ "$ENABLE_SSL_ON_NGINX" = "true" ] || [ "$ENABLE_SSL_ON_NGINX" = "1" ] || [ "$ENABLE_SSL_ON_NGINX" = "yes" ]; then
    echo "SSL enabled on nginx"
    export NGINX_LISTEN_PORT="443"
    export NGINX_SSL_CONFIG=" ssl"
    export NGINX_SSL_CERT="ssl_certificate /etc/nginx/ssl/cert.pem;"
    export NGINX_SSL_KEY="ssl_certificate_key /etc/nginx/ssl/key.pem;"
    export NGINX_SSL_PROTOCOLS="ssl_protocols TLSv1.2 TLSv1.3;"
    export NGINX_SSL_CIPHERS="ssl_ciphers HIGH:!aNULL:!MD5;"
else
    echo "SSL disabled on nginx"
    export NGINX_LISTEN_PORT="80"
    export NGINX_SSL_CONFIG=""
    export NGINX_SSL_CERT="# SSL disabled"
    export NGINX_SSL_KEY="# SSL disabled"
    export NGINX_SSL_PROTOCOLS="# SSL disabled"
    export NGINX_SSL_CIPHERS="# SSL disabled"
fi

# Generate nginx configuration from template
envsubst '${NGINX_LISTEN_PORT} ${NGINX_SSL_CONFIG} ${NGINX_SSL_CERT} ${NGINX_SSL_KEY} ${NGINX_SSL_PROTOCOLS} ${NGINX_SSL_CIPHERS}' < /etc/nginx/templates/default.conf.template > /etc/nginx/conf.d/default.conf

echo "Generated nginx configuration:"
cat /etc/nginx/conf.d/default.conf

# Execute nginx
exec nginx -g 'daemon off;'
