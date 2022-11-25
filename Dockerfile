FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY .env .
COPY --from=builder /app/server .
EXPOSE 8080
CMD [ "/app/server" ]
