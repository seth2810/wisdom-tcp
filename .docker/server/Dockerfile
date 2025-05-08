# Use the official Go image as the base image
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache make
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

# Use a smaller base image for the final image
FROM busybox
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 8080
CMD ["./server"]
