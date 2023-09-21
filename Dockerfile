FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src
RUN apk update
RUN apk upgrade

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./bin/inn-service cmd/api/main.go

FROM alpine AS runner

RUN apk --no-cache add bash make gcc musl-dev curl

COPY --from=builder /usr/local/src/bin/inn-service /

EXPOSE 8081

CMD /inn-service