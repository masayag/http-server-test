# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
#COPY go.sum go.sum

#COPY vendor/ vendor/

# Copy the go source
COPY cmd/main.go main.go
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o http_server_test main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/http_server_test .
USER 65532:65532

ENTRYPOINT ["/http_server_test"]
