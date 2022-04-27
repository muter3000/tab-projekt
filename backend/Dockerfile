############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/server ./cmd/server.go
############################
# STEP 2 build a small image
############################

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/server /go/bin/server
# Run the hello binary.
ENTRYPOINT ["/go/bin/server"]