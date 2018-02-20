FROM golang:latest
COPY . /go/src/github.com/haotianli89/driversvc
WORKDIR /go/src/github.com/haotianli89/driversvc/svc

RUN go build -o main main.go
RUN ls -la

CMD ["./main", "--registry=mdns"]