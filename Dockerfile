FROM golang:1.20.2-alpine3.16
WORKDIR /go/static_auth/
COPY ./src/ ./
RUN go mod tidy
RUN go build -o app .

FROM alpine:3.16
WORKDIR /static_auth/
COPY --from=0 /go/static_auth/app ./
CMD ["./app"]
