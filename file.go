package gox

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func FetchFile(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot new request")
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get request")
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, err
}
