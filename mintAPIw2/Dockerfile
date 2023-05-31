FROM golang:1.20-alpine as builder

ENV GO111MODULE=on
ENV GOPRIVATE=bitbucket.org/rappinc
ARG SSH_PRIVATE_KEY

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -tags musl -installsuffix cgo -o main ./cmd/api/

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
COPY --from=builder /app .
CMD ["./main"]
