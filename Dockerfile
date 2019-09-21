FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY . .

RUN cp -R /app/proto /go/src
RUN cp -R /app/store /go/src

RUN go get -v
RUN go build -o main .

CMD ["/app/main"]


#FROM alpine:latest
#
#RUN mkdir /app
#
#ADD mango /app
#
#WORKDIR /app
#
#CMD ["/app/mango"]
#
