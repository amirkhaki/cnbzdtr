FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o cdbot 
FROM alpine
WORKDIR /build
COPY --from=builder /build/cdbot /build/cdbot
EXPOSE 8080
CMD ["./cdbot"]
