package authentication

import (
	"BookServer/data"
	"encoding/json"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var cred data.Credential
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cr, ok := data.Creds[cred.Username]
	if !ok {
		http.Error(w, "Username not found", http.StatusNotFound)
		return
	}
	if cr.Password != cred.Password {
		http.Error(w, "Password mismatch", http.StatusUnauthorized)
		return
	}
	et := time.Now().Add(time.Hour * 24)
	claims := map[string]interface{}{
		"username": "suman",
		"exp":      et.Unix(),
	}
	_, tokenString, err := data.TokenAuth.Encode(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: et,
	})
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Successfully Logged In"))
	if err != nil {
		http.Error(w, "Can not Write Data", http.StatusInternalServerError)
		return
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Successfully Logged Out"))
	if err != nil {
		http.Error(w, "Can not Write Data", http.StatusInternalServerError)
		return
	}
}
