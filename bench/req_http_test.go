package bench

import (
	"strings"
	"testing"

	"github.com/AdityaMahaddalkar/go-http-benchmarks/util"
	"github.com/sirupsen/logrus"
)

func BenchmarkReqGet(b *testing.B) {

	for n := 0; n < b.N; n++ {
		for _, url := range util.UrlList {
			msg := ReqGet(url)
			logrus.Debug(msg)
		}
	}
}

func BenchmarkReqGetConcurrent(b *testing.B) {
	for n := 0;n < b.N;n ++ {

		ch := make(chan string)

		for _, url := range util.UrlList {
			go ReqGetConcurrent(url, ch)
		}

		for range util.UrlList {
			msg := <-ch
			logrus.Debug(msg)
		}
	}
}

func TestReqGet(t *testing.T) {

	for _, url := range util.UrlList {
		msg := ReqGet(url)
		logrus.Debug(msg)
		if !strings.Contains(msg, "Success") {
			t.Error(msg)
		}
	}
}