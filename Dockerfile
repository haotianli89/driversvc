FROM golang:latest
COPY . /go/src/github.com/haotianli89/driversvc
WORKDIR /go/src/github.com/haotianli89/driversvc/svc

RUN go get github.com/gocql/gocql
RUN go get github.com/golang/protobuf/proto
RUN go get github.com/micro/go-micro
RUN go get golang.org/x/net/context

RUN go build -o main main.go
RUN ls -la

CMD ["./main"]