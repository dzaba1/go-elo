FROM node:latest as builder

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build --prod

FROM nginx:latest

COPY --from=builder /app/dist/go-elo /usr/share/nginx/html

EXPOSE 80