package bench

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/AdityaMahaddalkar/go-http-benchmarks/util"
	"github.com/imroc/req/v3"
)


const (
	REQUEST_CREATION_ERROR_FORMAT = "Error in creating request object for %s: %v"
)

var (
	client = req.C()
)

func ReqGet(url string) string {
	resp, err := client.R().Get(url)
	if err != nil {
		return fmt.Sprintf(REQUEST_CREATION_ERROR_FORMAT, url, err)

	}
	defer resp.Body.Close()

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return fmt.Sprintf(util.RESPONSE_READING_ERROR_MESSAGE_FORMAT, url, err)
	}

	return fmt.Sprintf(util.RESPONSE_READING_SUCCESS_MESSAGE_FORMAT, nBytes, url)
}

func ReqGetConcurrent(url string, ch chan<- string) {

	resp, err := client.R().Get(url)
	if err != nil {
		ch <- fmt.Sprintf(REQUEST_CREATION_ERROR_FORMAT, url, err)
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
