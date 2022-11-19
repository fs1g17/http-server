FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /server
EXPOSE 9090
CMD ["/server"]
