FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goelo

FROM alpine

WORKDIR /app

COPY --from=builder /app/goelo /app/
COPY --from=builder /app/db/matches.csv /app/db/

ENV PORT=8080
ENV GIN_MODE=release
ENV ELO_KFACTOR=32
ENV ELO_INIT_RATING=1000
ENV ELO_LEAGUE_NAME="My League"

VOLUME /app/db/

ENTRYPOINT ["/app/goelo"]