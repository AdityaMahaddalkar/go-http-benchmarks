package bench

import (
	"strings"
	"testing"

	"github.com/AdityaMahaddalkar/go-http-benchmarks/util"
	"github.com/sirupsen/logrus"
)

func BenchmarkHeimdallHttpGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, url := range util.UrlList {
			msg := HeimdallHttpGet(url)
			logrus.Debug(msg)
		}
	}
}

func BenchmarkHeimdallHttpGetConcurrent(b *testing.B) {
	for n := 0;n < b.N;n ++ {

		ch := make(chan string)

		for _, url := range util.UrlList {
			go HeimdallHttpGetConcurrent(url, ch)
		}

		for range util.UrlList {
			msg := <-ch
			logrus.Debug(msg)
		}
	}
}

func TestHeimdallHttpGet(t *testing.T) {

	for _, url := range util.UrlList {
		msg := BaseHttpGet(url)
		logrus.Debug(msg)
		if !strings.Contains(msg, "Success") {
			t.Error(msg)
		}
	}
}

func TestHeimdallHttpGetConcurrent(t *testing.T) {

	ch := make(chan string)

	for _, url := range util.UrlList {
		go BaseHttpGetConcurrent(url, ch)
	}

	for range util.UrlList {
		msg := <-ch
		logrus.Debug(msg)
		if !strings.Contains(msg, "Success") {
			t.Error(msg)
		}
	}
}
