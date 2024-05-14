package gfunc

import (
	"encoding/json"
	"errors"
	"reflect"
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

func (q *Query) FindBy(by, field string, dataList []interface{}, Selected interface{}) error {

	err := q.Find(&dataList)
	if err != nil {
		return err
	}

	var all []interface{}
	all = append(all, dataList...)

	for _, v := range all {
		m, err := json.Marshal(v)
		if err != nil {
			return err
		}

		var tempMap map[string]interface{}
		if err := json.Unmarshal(m, &tempMap); err != nil {
			return err
		}

		value, found := tempMap[field]
		if !found {
			continue
		}

		if value == by {
			reflect.ValueOf(Selected).Elem().Set(reflect.ValueOf(v))
			return nil
		}
	}

	return errors.New("data not found")
}

func (q *Query) FindAllBy(by, field string, dataList []interface{}, Selected *[]interface{}) error {
	err := q.Find(&dataList)
	if err != nil {
		return err
	}

	var all []interface{}
	all = append(all, dataList...)

	for _, v := range all {
		m, err := json.Marshal(v)
		if err != nil {
			return err
		}

		var tempMap map[string]interface{}
		if err := json.Unmarshal(m, &tempMap); err != nil {
			return err
		}

		value, found := tempMap[field]
		if !found {
			continue
		}

		if value == by {
			*Selected = append(*Selected, v)
		}
	}

	if len(*Selected) == 0 {
		return errors.New("data not found")
	}

	return nil
}
