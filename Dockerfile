FROM golang:alpine as builder

WORKDIR /usr/src/app
COPY . .

ENV CGO_ENABLED=0 \
    GOOS=darwin \
    GOARCH=amr64

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

RUN go build -o main .

CMD ["./main"]
