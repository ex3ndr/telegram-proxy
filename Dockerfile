# build stage
FROM golang:1.9.1 AS build-env
RUN go get -d -v github.com/armon/go-socks5
ADD . /src
RUN cd /src && go build -ldflags "-linkmode external -extldflags -static" -o proxy

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-env /src/proxy /app/
CMD ["./proxy"]