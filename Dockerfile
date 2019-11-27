# Run command below to build binary.
#   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o main main.go
FROM golang:alpine AS build-env
RUN apk update
RUN apk add --no-cache git
WORKDIR /go/src/github.com/moodi/gateway-service
COPY . /go/src/github.com/moodi/gateway-service
RUN cd /go/src/github.com/moodi/gateway-service && go get -d -v
RUN go build -o ./main
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/moodi/gateway-service/main /app/
ENTRYPOINT [ "./main" ]
EXPOSE 80