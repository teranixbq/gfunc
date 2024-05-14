package gfunc

import (
	"errors"
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
