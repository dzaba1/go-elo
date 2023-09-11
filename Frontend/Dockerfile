FROM node:20.5.1 as builder

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build --prod

FROM nginx:latest

ENV API_URL=http://localhost:8080/

COPY --from=builder /app/dist/go-elo /usr/share/nginx/html

EXPOSE 80

CMD ["/bin/sh",  "-c",  "envsubst < /usr/share/nginx/html/assets/env.source.js > /usr/share/nginx/html/assets/env.js && exec nginx -g 'daemon off;'"]