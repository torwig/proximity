FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN make build

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/proxx .

CMD [ "/app/proxx" ]
