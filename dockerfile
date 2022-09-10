FROM golang:1.19.0-alpine3.16 as base
RUN apk add --no-cache gcc libc-dev make git

FROM base as builder
WORKDIR /app
RUN go env -w GOPROXY=direct
COPY go.sum .
COPY go.mod .
RUN go mod download
COPY . .
RUN make test
RUN make build-api
RUN make build-processor

FROM scratch as api
COPY --from=builder /app/api  /var/task/api
RUN chmod +x /var/task/api
CMD ["/var/task/api"]

FROM scratch as processor
COPY --from=builder /app/processor  /var/task/processor
RUN chmod +x /var/task/processor
CMD ["/var/task/processor"]