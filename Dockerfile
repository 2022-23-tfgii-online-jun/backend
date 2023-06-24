FROM golang:alpine AS builder

# Add dependencies
RUN apk add git
RUN apk --no-cache add tzdata
RUN apk add --no-cache --upgrade ffmpeg

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy only the necessary files for building
COPY cmd/api/main.go ./cmd/api/
COPY cmd/worker/main.go ./cmd/worker/
COPY internal/infra/api ./internal/infra/api
COPY internal/infra/worker ./internal/infra/worker
COPY internal/infra/repositories/postgresql ./internal/infra/repositories/postgresql
COPY config ./config

# Copy the Swagger documentation
COPY docs /docs

# Build the application
RUN go build -o /api ./cmd/api/main.go
RUN go build -o /worker ./cmd/worker/main.go

# Build a small image
FROM alpine

# Install FFmpeg
RUN apk add --no-cache --upgrade ffmpeg

# Copy from builder
COPY --from=builder /api /api
COPY --from=builder /worker /worker
COPY --from=builder /build/prod.env /
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /docs /docs

# Set Timezone
ENV TZ="America/Montevideo"

# Expose port out container
EXPOSE 8080

# Command to run both API and worker
CMD ["/bin/sh", "-c", "/api & /worker"]
