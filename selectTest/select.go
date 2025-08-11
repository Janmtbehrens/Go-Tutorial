package selectTest

import (
	"net/http"
	"time"
)

func MeasureURL(url string) (string, time.Duration) {
	start := time.Now()
	http.Get(url)
	dur := time.Since(start)

	return url, dur
}

func Racer(urlA, urlB string) (winner string) {
	select {
	case <-ping(urlA):
		return urlA
	case <-ping(urlB):
		return urlB
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
