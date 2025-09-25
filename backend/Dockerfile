# backend/Dockerfile
FROM golang:1.22-alpine AS build
WORKDIR /app

# 1) Salin go.mod/go.sum terlebih dulu agar caching dependency tetap optimal
COPY go.mod go.sum ./
RUN go mod download

# 2) Salin seluruh source, lalu tidy supaya go.sum terisi sesuai import aktual
COPY . .
RUN go mod tidy

# 3) Build
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# --- runtime image ---
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/server /app/server
ENV PORT=8080
EXPOSE 8080
CMD ["/app/server"]
