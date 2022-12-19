FROM golang:1.17 AS build

ADD . /app
WORKDIR /app
RUN go build ./cmd/main.go

FROM ubuntu:20.04

COPY . .
COPY --from=build /app/main/ .

CMD ./scripts/gen_ca.sh
CMD ./scripts/gen_cert.sh mail.ru 1000000000000

EXPOSE 8080

CMD apt-get install ca-certificates -y
ADD certs/ca.crt /usr/local/share/ca-certificates/ca.crt
CMD chmod 644 /usr/local/share/ca-certificates/ca.crt && update-ca-certificates

CMD ./main