############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/server ./cmd/server/server.go
############################
# STEP 2 build a small image
############################

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/server /server
COPY ./swagger.yaml /swagger.yaml
# Run the hello binary.
ENTRYPOINT ["/server"]