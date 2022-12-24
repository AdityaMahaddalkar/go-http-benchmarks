package bench

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/AdityaMahaddalkar/go-http-benchmarks/util"
)

///BaseHttpGet will execute HTTP GET method with url
///and return the number of bytes received
func BaseHttpGet(url string) string {

	resp, err := http.Get(url)
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

///BaseHttpGetConcurrent will execute HTTP GET method with url
///and send the number of bytes to channel ch
func BaseHttpGetConcurrent(url string, ch chan<- string) {

	resp, err := http.Get(url)
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
