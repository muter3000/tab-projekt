############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/auth ./cmd/auth/auth.go
############################
# STEP 2 build a small image
############################

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/auth /go/bin/auth
# Run the hello binary.
ENTRYPOINT ["/go/bin/auth"]