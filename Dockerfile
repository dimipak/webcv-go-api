FROM golang:latest AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get ./...

RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o admin main.go

# Use distroless base image without providing CGO_ENABLED=0 flag on build so the binary file will use c library from host
# machine and it will be faster. Otherwise use scratch like below
#FROM gcr.io/distroless/base-debian11 AS release

# For scratch based image we need to use CGO_ENABLED=0 when building image
FROM scratch AS release

WORKDIR /

COPY --from=builder /usr/src/app/admin .
COPY --from=builder /usr/src/app/.env .


ENTRYPOINT ["/admin"]