FROM golang:1.10
ADD . .
CMD go run main.go
EXPOSE 9000

