FROM golang:1.11.4 as builder
WORKDIR /todos
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /todos .
CMD ["./todos"]