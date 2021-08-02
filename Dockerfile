FROM golang:1.16.5
MAINTAINER rafaelmedrado
WORKDIR app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o api main.go
ENTRYPOINT ./api
EXPOSE 5000
