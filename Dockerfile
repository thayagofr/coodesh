FROM golang:alpine
RUN mkdir /app
ADD ./http /app
WORKDIR /app
RUN go mod download
RUN go build main.go
EXPOSE 8080
CMD ["./main"]