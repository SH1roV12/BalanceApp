package middleware

import (
	"log"
	"net"
	"net/http"
	"strings"
)




func IPFilterMiddleware(next http.Handler, allowedCIDRs []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteIP := r.RemoteAddr
		log.Printf("Initial remote IP: %s", remoteIP)

		// Проверяем X-Real-IP, если доступен.
		if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
			log.Printf("Using X-Real-IP header: %s", realIP)
			remoteIP = realIP
		}

		// Разделяем адрес на хост и порт.
		var host string
		if strings.Contains(remoteIP, ":") {
			var err error
			host, _, err = net.SplitHostPort(remoteIP)
			if err != nil {
				http.Error(w, "Invalid remote IP address", http.StatusBadRequest)
				return
			}
		} else {
			host = remoteIP
		}

		// Проверяем, разрешен ли IP-адрес.
		if !IsIPAllowed(host, allowedCIDRs) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Передаем управление дальше.
		next.ServeHTTP(w, r)
	})
}

// Проверяет, входит ли IP-адрес в разрешенные диапазоны.
func IsIPAllowed(ip string, allowedCIDRs []string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, cidr := range allowedCIDRs {
		_, allowedNet, err := net.ParseCIDR(cidr)
		if err != nil {
			continue
		}
		if allowedNet.Contains(parsedIP) {
			return true
		}
	}
	return false
}