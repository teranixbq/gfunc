package gfunc

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

type Query struct {
	jsonfile []byte
}

func NewJsonFile(file string) (*Query, error) {
	if !strings.HasSuffix(file, ".json") {
		return nil, errors.New("file must be json")
	}

	bytefile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	b := Query{
		jsonfile: bytefile,
	}

	return &b, nil
}

func NewJsonUrl(url string) (*Query, error) {
	if !strings.HasSuffix(url, ".json") {
		return nil, errors.New("URL must point to a JSON file")
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch JSON file: " + response.Status)
	}

	bytefile, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	b := Query{
		jsonfile: bytefile,
	}

	return &b, nil
}