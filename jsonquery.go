package gfunc

import (
	"encoding/json"
	"errors"
)

func (q *Query) Find(data interface{}) error {
	errPointer := CheckStructPointer(data)
	if errPointer != nil {
		return errPointer
	}

	err := json.Unmarshal(q.jsonfile, data)
	if err != nil {
		return err
	}

	return nil
}

// Parameter dataList is []interface{}, so you need to use for loop to append the data e.g :
//
//	findData := []data{}
//	var datalist []interface{}
//
//	q, errJson := NewJsonFile("your_json_file.json")
//	for _, item := range findData {
//		datalist = append(datalist, item)
//	}
//
//	result,err := q.FindBy("3", "id", datalist)
func (q *Query) FindBy(by, field string, dataList []interface{}) (interface{}, error) {
	var Selected interface{}
	err := q.Find(&dataList)
	if err != nil {
		return nil, err
	}

	var all []interface{}
	all = append(all, dataList...)

	for _, v := range all {
		m, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		var tempMap map[string]interface{}
		if err := json.Unmarshal(m, &tempMap); err != nil {
			return nil, err
		}

		value, found := tempMap[field]
		if !found {
			continue
		}

		// if value == by {
		// 	reflect.ValueOf(Selected).Elem().Set(reflect.ValueOf(v))
		// 	return nil,nil
		// }

		if value == by {
			Selected = v
			break
		}

	}

	return Selected, nil
}

func (q *Query) FindAllBy(by, field string, dataList []interface{}) ([]interface{}, error) {
	var Selected []interface{}
	err := q.Find(&dataList)
	if err != nil {
		return nil, err
	}

	var all []interface{}
	all = append(all, dataList...)

	for _, v := range all {
		m, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		var tempMap map[string]interface{}
		if err := json.Unmarshal(m, &tempMap); err != nil {
			return nil, err
		}

		value, found := tempMap[field]
		if !found {
			continue
		}

		if value == by {
			Selected = append(Selected, v)
		}
	}

	if len(Selected) == 0 {
		return nil, errors.New("data not found")
	}

	return Selected, nil
}
