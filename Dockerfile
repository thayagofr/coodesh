FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build http/main.go
EXPOSE 8080
CMD ["/app/main"]