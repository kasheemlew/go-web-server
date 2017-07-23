package setCookie

import (
	"net/http"
	"time"
)

func main() {
	// Setting Cookie
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name:"username", Value:"kasheemlew", Expires:expiration}
	http.SetCookie(w, cookie)

	// Fetch Cookie
	cookie, _ := r.Cookie("username")
}