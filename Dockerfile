FROM golang:1.13
COPY . /src/github.com/lazywhite/mariadb-broker
WORKDIR /src/github.com/lazywhite/mariadb-broker
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /mariadb-broker . && strip /mariadb-broker

FROM alpine:3.6
COPY --from=0 /mariadb-broker /mariadb-broker
CMD ["/mariadb-broker", "-logtostderr"]
