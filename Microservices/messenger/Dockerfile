FROM golang:1.20

WORKDIR /app

COPY source_code/ .

RUN go mod init messenger && go mod tidy || true

RUN go build -o main .

EXPOSE 3002

# Build number shown on the home page (set by CI with --build-arg)
ARG APP_VERSION=dev
ENV APP_VERSION=$APP_VERSION

CMD ["./main"]
