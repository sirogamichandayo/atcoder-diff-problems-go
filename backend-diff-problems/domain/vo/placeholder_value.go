package vo

import "fmt"

type PlaceholderValue struct {
	placeholder string
	values      []interface{}
}

func (pv PlaceholderValue) Placeholder() string {
	return pv.placeholder
}

func (pv PlaceholderValue) Values() []interface{} {
	return pv.values
}

type PlaceholderValueList []PlaceholderValue

const MaxPlaceHolderValueLen = 1000

func NewPlaceHolderValue(placeholder string, values []interface{}) (PlaceholderValue, error) {
	if len(values) > MaxPlaceHolderValueLen {
		return PlaceholderValue{}, fmt.Errorf("value length %d exceeds max value length %d", len(values), MaxPlaceHolderValueLen)
	}
	return PlaceholderValue{placeholder: placeholder, values: values}, nil
}
