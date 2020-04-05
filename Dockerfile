FROM golang:1.13-alpine AS build

ENV GO111MODULE on

RUN addgroup -S scratchuser && adduser -S scratchuser -G scratchuser -u 12345
RUN apk --no-cache --update add ca-certificates && update-ca-certificates

ADD . /go/src/guthub.com/baba2k/graphql-crud-server/
WORKDIR /go/src/guthub.com/baba2k/graphql-crud-server/server/docker

RUN go get -d -v -t ./... \
        && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM scratch

ADD $SCHEMA_FILE $SCHEMA_FILE
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/guthub.com/baba2k/graphql-crud-server/server/docker/app /bin/app

USER scratchuser
EXPOSE 8080

CMD ["/bin/app"]