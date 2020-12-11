FROM node:14.15.0-buster-slim as builder

WORKDIR /app
COPY package.json package-lock.json ./

RUN npm ci

COPY . .
RUN npm run build -- --configuration=production,no-server


FROM nginx:1.19.3 as runtime

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist/doogeun-web-front /usr/share/nginx/html
