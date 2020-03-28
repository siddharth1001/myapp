FROM golang:1.13 as builder

LABEL maintainer="siddharth.rawat1001@gmail.com"

WORKDIR /mygoapp

COPY go.mod go.mod
COPY go.sum go.sum

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/base-debian10

FROM golang:1.13

WORKDIR /

COPY --from=builder /mygoapp/myapp .
COPY --from=builder /mygoapp/assets ./assets

# EXPOSE 8090 8090

# CMD ["/myapp"]
ENTRYPOINT ["/myapp"]
