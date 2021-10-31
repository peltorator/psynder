# Build
FROM golang:1.17.2-alpine3.14 as builder
ADD . /build
WORKDIR /build
RUN go build -o server github.com/peltorator/psynder/cmd/server

# Deploy
FROM alpine:3.14
EXPOSE 443
COPY --from=builder /build/server /
ENTRYPOINT [ "/server" ]
