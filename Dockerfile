FROM golang:1.22 as builder

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Install Node.js and npm
RUN apk add --update nodejs npm
# Install javascript-obfuscator globally
RUN npm install -g javascript-obfuscator

WORKDIR /root/
COPY --from=builder /app/main .
# Ensure the entire public directory is copied into the Docker image
COPY --from=builder /app/public ./public
COPY --from=builder /app/scripts ./scripts
COPY --from=builder /app/.env .env
COPY --from=builder /app/cert ./cert

# Make sure the img directory exists in the image
RUN mkdir -p /root/public/img
# Copy the images into the img directory
COPY --from=builder /public/img ./public/img

EXPOSE 80 443 587

CMD ["./main"]
