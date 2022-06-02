# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# # Add Maintainer info
# LABEL maintainer="Steven Victor <chikodi543@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .


#Builder image build the go binary : Setting the alias as builder
# FROM golang:1.15-alpine as builder
# RUN mkdir /app

# # add this file directory to the image
# ADD . /app
# # directory where the source file execution takes place
# WORKDIR /app

# # RUN some go commands
# RUN go clean --modcache
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/

# # Our production image used to run our app
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# RUN apk add --no-cache git make musl-dev go
# COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

# # Add docker-compose-wait tool -------------------
# ENV DOCKER_ENV=/root/.env
# ENV WAIT_VERSION 2.7.2
# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
# RUN chmod +x /wait


# # Configure GO
# ENV GOROOT /usr/lib/go
# ENV GOPATH /go
# ENV PATH /go/bin:$PATH

# RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
EXPOSE 8080
CMD ["./main"]