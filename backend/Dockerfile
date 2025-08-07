FROM golang:1.23-alpine as builder

WORKDIR /app
RUN apk add 
RUN nslookup dl-cdn.alpinelinux.org
RUN ping -c 4 dl-cdn.alpinelinux.org
RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . ./

RUN go build -o ./bin/cmd cmd/main.go

FROM alpine:latest AS runner

COPY --from=builder /app/bin/cmd ./

CMD ["sh", "-c", "until nc -z db 5432; do sleep 1; done && ./cmd"]
