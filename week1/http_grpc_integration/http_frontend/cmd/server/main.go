package main

import "time"

const (
	// Адреса серверов
	grpcBackendAddress = "localhost:50051"
	httpAddress        = ":8080"

	// HTTP таймауты
	httpReadHeaderTimeout = 5 * time.Second  // Защита от Slowloris атаки
	httpReadTimeout       = 15 * time.Second // Лимит на чтение всего запроса
	httpWriteTimeout      = 15 * time.Second // Лимит на запись ответа
	httpIdleTimeout       = 60 * time.Second // Таймаут keep-alive соединений
	httpShutdownTimeout   = 5 * time.Second  // Таймаут graceful shutdown

	// gRPC клиент keepalive параметры
	grpcKeepaliveTime    = 5 * time.Minute // Интервал ping'ов для обнаружения мёртвого сервера
	grpcKeepaliveTimeout = 1 * time.Second // Таймаут ожидания pong
)

func main() {

}
