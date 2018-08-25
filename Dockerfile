FROM golang:latest
RUN mkdir /goa
ADD . /goa
WORKDIR /goa

RUN go get github.com/gin-gonic/gin

RUN go build -o main .

EXPOSE 8080


