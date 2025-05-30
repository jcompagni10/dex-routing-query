package skip

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	baseURL = "https://api.skip.build/v2"
)

var APIKey = os.Getenv("SKIP_API_KEY")

func PostRequest(path string, reqBody io.Reader) ([]byte, error) {
	url := baseURL + path

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", APIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(body))
	}

	return body, err
}

func GetRequest(path string, queryParams url.Values) ([]byte, error) {
	baseUrl := baseURL + path
	if len(queryParams) > 0 {
		baseUrl += "?" + queryParams.Encode()
	}

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	return body, err
}
