# Flutter Web
FROM ghcr.io/cirruslabs/flutter:3.41.7 AS builder

ENV FLUTTER_ALLOW_HTTP=true
USER root

WORKDIR /app
COPY . .

RUN flutter pub get
RUN flutter build web --release

FROM nginx:alpine
COPY --from=builder /app/build/web /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
