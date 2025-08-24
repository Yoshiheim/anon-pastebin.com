package helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mileusna/useragent"
)

func MiddlewareHand(urla string, hand http.Handler) {
	http.Handle(urla, LoggingMiddleware(hand))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s - %s", r.Method, r.URL.Path, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}

func IsChrome(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		if ua.OS == useragent.Chrome && ua.Desktop {
			log.Println("okay, its chrome and desktop!")
			next.ServeHTTP(w, r)
			return
		}
		log.Println("its not chrome and desktop, haha")
		fmt.Fprintln(w, "chrome and desktop required, haha")
	})
}

func IsFireFox(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		if ua.OS == useragent.Firefox {
			log.Println("okay, its FireFox and Desktop!")
			next.ServeHTTP(w, r)
			return
		}
		log.Println("its not FireFox and Desktop, haha")
		fmt.Fprintln(w, "Firefox and Desktop required, haha")
	})
}

// check is userAgent uses Linux lmao
func IsLinux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		if ua.OS == useragent.Linux {
			log.Println("okay, this bro use linux(i hope its Arch btw)")
			next.ServeHTTP(w, r)
			return
		}
		log.Println("not linux...")
		fmt.Fprintf(w, "bro, use linux, windows is shit(Arch better linux distro btw)")
	})
}
