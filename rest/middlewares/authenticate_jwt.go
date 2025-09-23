package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	//"ecommerce/repo"
	"encoding/base64"
	//"encoding/json"
	//"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	headerArr := strings.Split(header, " ")

	if len(headerArr) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	accessToken := headerArr[1]

	tockenParts := strings.Split(accessToken, ".")
	if len(tockenParts) != 3 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtHeader := tockenParts[0]
	jwtPayload := tockenParts[1]
	signature := tockenParts[2]

	message := jwtHeader + "." + jwtPayload
	
	byteArrSecret := []byte(m.cnf.JwtSecretKey)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	hash := h.Sum(nil)
	newSignature := base64UrlEncode(hash)

	if signature != newSignature {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	next.ServeHTTP(w, r)
 })
		
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}