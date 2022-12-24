package bench

import (
	"strings"
	"testing"

	"github.com/AdityaMahaddalkar/go-http-benchmarks/util"
	"github.com/sirupsen/logrus"
)

func BenchmarkBaseHttpGet(b *testing.B) {

	for n := 0; n < b.N; n++ {
		for _, url := range util.UrlList {
			msg := BaseHttpGet(url)
			logrus.Debug(msg)
		}
	}
}

func BenchmarkTestBaseHttpGetConcurren(b *testing.B) {
	for n := 0;n < b.N;n ++ {

		ch := make(chan string)

		for _, url := range util.UrlList {
			go BaseHttpGetConcurrent(url, ch)
		}

		for range util.UrlList {
			msg := <-ch
			logrus.Debug(msg)
		}
	}
}

func TestBaseHttpGet(t *testing.T) {

	for _, url := range util.UrlList {
		msg := BaseHttpGet(url)
		logrus.Debug(msg)
		if !strings.Contains(msg, "Success") {
			t.Error(msg)
		}
	}
}

func TestBaseHttpGetConcurrent(t *testing.T) {

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
