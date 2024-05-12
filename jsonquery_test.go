package gfunc

import (
	"log"
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

	q, errJson := NewJsonFile("example.json")
	err := q.Find(&listData)

	assert.NoError(t, errJson)
	assert.NoError(t, err)
	assert.NotNil(t, listData)
	assert.NotEmpty(t, listData)
}

func TestFindBy(t *testing.T) {
	findData := []data{}
	var datalist []interface{}
	var selectedData interface{}
	

	q, errJson := NewJsonFile("example.json")
	for _, item := range findData {
		datalist = append(datalist, item)
	}

	err := q.FindBy("3", "id", datalist, &selectedData)

	assert.NoError(t, errJson)
	assert.NoError(t, err)
	assert.NotNil(t, selectedData)
	assert.NotEmpty(t, selectedData)
	
}

func TestFindAllBy(t *testing.T) {
	findData := []data{}
	var datalist []interface{}
	var selectedData []interface{}
	

	q, errJson := NewJsonFile("example.json")
	for _, item := range findData {
		datalist = append(datalist, item)
	}
	err := q.FindAllBy("Fruit", "category", datalist,&selectedData)

	assert.NoError(t, errJson)
	assert.NoError(t, err)
	assert.NotNil(t, findData)
	log.Println(selectedData)
}