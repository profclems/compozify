FROM node:18-alpine AS webBuilder
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY ./web ./
RUN rm -rf ./node_modules
RUN yarn global add pnpm && pnpm install && pnpm build

FROM golang:1.20 AS builder
WORKDIR /app/compozify-web
COPY . .
COPY --from=webBuilder /app/dist ./web/dist
RUN go build -o ./bin/compozify-web ./cmd/compozify-web

FROM alpine
RUN apk add --no-cache libc6-compat
RUN mkdir /app
EXPOSE 8080
COPY --from=builder /app/compozify-web/bin/compozify-web /app/compozify-web
WORKDIR /app
COPY cmd/compozify-web/entrypoint /entrypoint
ENTRYPOINT ["/entrypoint"]