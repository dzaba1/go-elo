FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goelo

FROM alpine

COPY --from=builder /app/goelo /
COPY --from=builder /app/db/matches.csv /db/

ENV PORT=8080
ENV GIN_MODE=release
ENV ELO_KFACTOR=32
ENV ELO_INIT_RATING=1000
ENV ELO_LEAGUE_NAME="My League"

EXPOSE 8080

ENTRYPOINT ["/goelo"]