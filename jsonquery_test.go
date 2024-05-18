package gfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func TestFind(t *testing.T) {
	listData := []data{}

	q, errJson := NewJsonFile("_example/example.json")
	err := q.Find(&listData)

	assert.NoError(t, errJson)
	assert.NoError(t, err)
	assert.NotNil(t, listData)
	assert.NotEmpty(t, listData)
}

func TestFindBy(t *testing.T) {
	findData := []data{}
	var datalist []interface{}

	q, errJson := NewJsonFile("_example/example.json")
	for _, item := range findData {
		datalist = append(datalist, item)
	}

	result, err := q.FindBy("3", "id", datalist)

	assert.NoError(t, errJson)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func TestFindAllBy(t *testing.T) {
	findData := []data{}
	var datalist []interface{}

	q, errJson := NewJsonFile("_example/example.json")
	for _, item := range findData {
		datalist = append(datalist, item)
	}
	result, err := q.FindAllBy("Fruit", "category", datalist)

	assert.NoError(t, errJson)
	assert.NoError(t, err)
	assert.NotNil(t, findData)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}
