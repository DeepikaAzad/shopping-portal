# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV APP_PORT=8080 \
    APP_ENV=development \
    APP_DEBUG=false \
    DB_CONNECTION=tcp \
    DB_HOST=host.docker.internal \
    DB_PORT=3306 \
    DB_DATABASE=sezzle \
    DB_USERNAME=root \
    DB_PASSWORD=root \
    SECRET=0YL45Z9HBj 


# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .
# Export necessary port
EXPOSE 8080
# Command to run when starting the container
CMD ["./main"]