############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/db ./cmd/create-database/create-database.go
############################
# STEP 2 build a small image
############################

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/db /go/bin/db
# Run the hello binary.
ENTRYPOINT ["/go/bin/db"]