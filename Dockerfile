
FROM golang:1.12.0-alpine3.9

#need for working git and github libr

RUN go mod download

RUN apk add --no-cache git

RUN mkdir /resttest

ADD . /resttest

WORKDIR /resttest

RUN go build -o main .

CMD ["/resttest/main"]

