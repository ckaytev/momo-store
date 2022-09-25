FROM nginx:stable-alpine
WORKDIR /app
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./public_html /usr/share/nginx/html
EXPOSE 80
