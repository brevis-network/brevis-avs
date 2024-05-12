# ================================
# Build image
# ================================
FROM golang:1.22.3-alpine3.19 as build

# Set up a build area
ARG BUILD_DIR=/app
RUN mkdir -p ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

# Copy working directory
COPY . .

# Build brevis
RUN go build -o brevis

# ================================
# Run image
# ================================
FROM alpine:3.19

# Create a brevis user and group with /app as its home directory
RUN addgroup -S brevis && adduser -k /dev/null -h /app -G brevis -S brevis

# Switch to the new home directory
WORKDIR /app

# Copy built executable and any staged resources from builder
COPY --from=build --chown=brevis:brevis /app/brevis /app/brevis

# Ensure all further commands run as the brevis user
USER brevis:brevis

# Start the brevis service when the image is run, default to listening on 8081 in production environment
ENTRYPOINT ["./brevis"]
CMD ["run"]
