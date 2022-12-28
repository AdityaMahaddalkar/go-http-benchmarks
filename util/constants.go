package util

const (
	GET_ERROR_MESSAGE_FORMAT                = "Error in fetching %s: %v"
	RESPONSE_READING_ERROR_MESSAGE_FORMAT   = "Error in reading response from %s: %v"
	RESPONSE_READING_SUCCESS_MESSAGE_FORMAT = "Successfully read %d bytes from %s"
)

var (
	UrlList = []string{
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://bbc.co.uk",
		"https://amazon.co.uk",
		"https://twitter.com",
		"https://wikipedia.org",
		"https://ebay.co.uk",
		"https://live.com",
	}
)
