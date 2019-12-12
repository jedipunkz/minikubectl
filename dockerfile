FROM golang:1.13

WORKDIR /go/src/minikubectl
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "/go/src/minikubectl/main.go"]
