# Start from the latest golang base image
FROM golang:alpine

RUN GOCACHE=OFF

RUN go env -w GOPRIVATE=github.com/StanDenisov


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

RUN apk add git

RUN git config --global url."https://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@github.com".insteadOf "https://github.com"

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 13200

#ENTRYPOINT ["/app"]

# Command to run the executable
CMD ["./main"]