FROM php:8.1-cli
WORKDIR /app
COPY source_code/ .
EXPOSE 3004
# Build number shown on the home page (set by CI with --build-arg)
ARG APP_VERSION=dev
ENV APP_VERSION=$APP_VERSION

CMD ["php", "-S", "0.0.0.0:3004", "index.php"]
