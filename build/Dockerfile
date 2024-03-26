FROM golang:1.20
LABEL maintainer=rafaelmedrado
WORKDIR app
COPY . .
RUN go mod tidy
RUN go build -o api cmd/main.go
ENTRYPOINT ./api