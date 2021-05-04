FROM golang:1.16-alpine3.13

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8085
CMD ["fresh"]
