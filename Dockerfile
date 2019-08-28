# Build the go artifact
FROM golang:latest as builder
LABEL maintainer="Clint Colding <clintcolding@gmail.com>"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create container using artifact from build stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8082
ENTRYPOINT ["./main"]