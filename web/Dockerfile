# Stage 1
FROM node:22-alpine3.20 as builder

WORKDIR /app
COPY . .
RUN npm ci && npm run build

# Stage 2
FROM node:22-alpine3.20

WORKDIR /app
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

CMD ["node", "server.js"]