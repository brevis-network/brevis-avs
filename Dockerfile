FROM --platform=linux/amd64 golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o brevis .

FROM --platform=linux/amd64 alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/brevis .

# Set the entrypoint command
ENTRYPOINT ["/app/brevis"]
