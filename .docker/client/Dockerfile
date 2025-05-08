# Use the official Go image as the base image
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache make
WORKDIR /app
COPY . .
RUN go build -o client cmd/client/main.go

# Use a smaller base image for the final image
FROM busybox
WORKDIR /app
COPY --from=builder /app/client /app/client
CMD ["./client"]
