# goenv: Build a Go environment
FROM golang:1.21.0-alpine as goenv

RUN apk update

WORKDIR /app

# Common resources
COPY go.mod .
COPY go.sum .
RUN go mod download

# Place server code
COPY server server/

# builder: Build server
FROM goenv as builder

# Set build tag
ARG build

# Reconfigure DI to match build tag
RUN go run github.com/google/wire/cmd/wire \
    gen -tags "$build" ./.../injection/...

# Build server
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/main ./server/cmd

# air: Prepare debug environment
FROM goenv as dev

# HotSwap library introduced
RUN go install github.com/cosmtrek/air@latest

# Set build tag
ARG build
ENV BUILD $build

# Reconfigure DI to match build tag
# (local environment will also change)
CMD go run github.com/google/wire/cmd/wire \
    gen -tags "$BUILD" ./.../injection/... && \
    air -c ./server/.air.toml

# run: Construction of runtime environment
FROM gcr.io/distroless/static-debian10 as run

COPY --from=builder /app/main /main
# If files are needed, place them in `etc` folder
# COPY --from=builder /app/etc/ /etc/

CMD ["/main"]
