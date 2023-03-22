FROM golang:1.20-alpine as builder
WORKDIR /src/app
COPY go.mod go.sum main.go ./
RUN go build -o server

FROM alpine
WORKDIR /root/
COPY --from=builder /src/app ./app
EXPOSE 8080
CMD ["./app/server"]