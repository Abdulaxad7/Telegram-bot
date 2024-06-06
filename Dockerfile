FROM golang:1.22.3 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM alpine:3.14 as final
WORKDIR /
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/main"]