FROM nginx:1.23.1-alpine
COPY docker/nginx/nginx.conf /etc/nginx/nginx.conf

COPY docker/nginx/default.conf.tpl /etc/nginx/conf.d/default.conf
ARG APP_WSS_PORT
RUN sed -i "s/__lport__/$APP_WSS_PORT/gI" /etc/nginx/conf.d/default.conf