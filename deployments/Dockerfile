FROM golang:1.26.6-alpine3.23 AS build

# install build tools
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
RUN apk update && apk add git && go env -w GOPROXY=direct

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY . ${ROOT}
WORKDIR ${ROOT}
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o $ROOT/main ./cmd/api/main.go

# copy artifacts to a clean image
FROM alpine:3.23 AS prod
ENV ROOT=/go/src/app
COPY --from=build ${ROOT}/main main
ENTRYPOINT ["/main"]
