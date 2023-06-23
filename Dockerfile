FROM golang:alpine AS builder

# Add dependencies
RUN apk add git
RUN apk --no-cache add tzdata
RUN apk add --no-cache --upgrade ffmpeg

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o ../cmd/api/main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp ../cmd/api/main .
RUN cp /build/dev.env .

# Build a small image
FROM alpine

# Install FFmpeg
RUN apk add --no-cache --upgrade ffmpeg

# Copy from builder
COPY --from=builder /dist/main /
COPY --from=builder /build/dev.env .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set Timezone
ENV TZ="America/Montevideo"

# Expose port out container
EXPOSE 8080

# Command to run
CMD ["/main"]
