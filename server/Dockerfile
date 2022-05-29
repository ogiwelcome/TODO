FROM golang:1.16.3-alpine

RUN apk add --update &&  apk add git curl

ADD ./app /go/src/app

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get -u github.com/cosmtrek/air

RUN go mod download

EXPOSE 8000

CMD ["air"]