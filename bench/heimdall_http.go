package bench

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/AdityaMahaddalkar/go-http-benchmarks/util"
	"github.com/gojek/heimdall/httpclient"
)

var heimdallClient = httpclient.NewClient(httpclient.WithHTTPTimeout(10 * time.Second))

func HeimdallHttpGet(url string) string {

	resp, err := heimdallClient.Get(url, nil)
	if err != nil {
		return fmt.Sprintf(util.GET_ERROR_MESSAGE_FORMAT, url, err)
	}
	defer resp.Body.Close()

	n_bytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return fmt.Sprintf(util.RESPONSE_READING_ERROR_MESSAGE_FORMAT, url, err)

	}

	return fmt.Sprintf(util.RESPONSE_READING_SUCCESS_MESSAGE_FORMAT, n_bytes, url)

}

func HeimdallHttpGetConcurrent(url string, ch chan<- string) {

	resp, err := heimdallClient.Get(url, nil)
	if err != nil {
		ch <- fmt.Sprintf(util.GET_ERROR_MESSAGE_FORMAT, url, err)
		return
	}
	defer resp.Body.Close()

	n_bytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf(util.RESPONSE_READING_ERROR_MESSAGE_FORMAT, url, err)
		return
	}

	ch <- fmt.Sprintf(util.RESPONSE_READING_SUCCESS_MESSAGE_FORMAT, n_bytes, url)
}