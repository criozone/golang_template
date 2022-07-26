server {
    listen 80;

    location /{
        proxy_set_header   X-Forwarded-For $remote_addr;
        proxy_set_header   Host $http_host;
        proxy_pass https://wss-go:__lport__;
    }
}
