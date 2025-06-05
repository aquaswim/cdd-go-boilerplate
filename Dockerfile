FROM golang:1.23-alpine AS builder-go
RUN apk add --no-cache make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM alpine:latest
RUN apk add --no-cache tzdata
# Set timezone to Asia/Jakarta
ENV TZ=Asia/Jakarta

# Create a non-root user
RUN addgroup -S app \
    && adduser -S app -G app

COPY --from=builder-go /app/oapi-server /usr/local/bin/oapi-server
EXPOSE 3000

USER app

ENTRYPOINT ["oapi-server"]