FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/bin/api ./app/bin/api
COPY --from=builder /app/utils/.env ./app/utils/.env
# See docker image structure: `docker run -it gojira-project-manager_app sh` and use 'ls -lsa'
EXPOSE 3000
CMD ./app/bin/api