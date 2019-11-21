# Run command below to build binary.
#   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o main main.go
FROM golang:alpine AS build-env
RUN apk update
RUN apk add --no-cache git
WORKDIR /src/gateway
COPY . .
RUN go get -d -v
RUN cd /src/gateway && go build -o main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /src/gateway/ /app/
ENTRYPOINT [ "./main" ]
EXPOSE 80