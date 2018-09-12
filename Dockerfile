FROM golang:latest
RUN mkdir /goa
ADD . /goa
WORKDIR /goa

RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get golang.org/x/crypto/bcrypt

RUN go build -o main .

EXPOSE 8080

CMD [ "/goa/main" ]


