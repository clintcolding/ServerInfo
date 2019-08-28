# ServerInfo

Simple Go application that returns basic server stats on port 8082.

### Dockerfile

The dockerfile has multiple stages, the first stage is used to compile the main.go file into an executable and the second stage builds the executable into a stripped down Alpine based container.