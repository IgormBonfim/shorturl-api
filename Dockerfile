FROM golang:1.23-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM scratch AS publish

COPY --from=base /app/main /

ENTRYPOINT [ "./main" ]