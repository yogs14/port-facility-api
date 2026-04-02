# Tahap 1: Membangun aplikasi
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
# Kompilasi aplikasi menjadi sebuah file binary bernama 'port-api'
RUN go build -o port-api main.go

# Tahap 2: Membuat container super ringan untuk produksi
FROM alpine:latest
WORKDIR /root/
# Ambil file binary dari Tahap 1
COPY --from=builder /app/port-api .
EXPOSE 8080
# Perintah saat container dinyalakan
CMD ["./port-api"]