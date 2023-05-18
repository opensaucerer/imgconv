package imgconv

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(url string) (io.ReadCloser, error) {
	errMsg := fmt.Sprintf("impexp can't find the file or file is invalid %s", url)
	if !isValidURL(url) {
		return nil, fmt.Errorf("impexp found a invalid URL: %s", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(errMsg)
	}
	return resp.Body, nil
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	return err == nil
}
