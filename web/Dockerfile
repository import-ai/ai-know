# Stage 1
FROM node:22 AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2
FROM nginx:1.27
COPY --from=builder /app/dist /usr/share/nginx/html
