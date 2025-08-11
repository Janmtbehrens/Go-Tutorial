package selectTest

import "time"
import "http"

func MeasureURL(url string) (string, int) {
	start := time.Now()
	http.Get(urlA)
	dur := time.Since(start)

	return url, dur
}

func Racer(urlA, urlB string) string {
	_, startA := MeasureURL(urlA)
	_, startB := MeasureURL(urlB)

	if(startA < startB) {
		return urlA
	}

	return urlB
}
