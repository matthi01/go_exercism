package flatten

func Flatten(nested interface{}) []interface{} {
	var flat = []interface{}{}
	switch nested.(type) {
	case []interface{}:
		nestedArr := nested.([]interface{})
		for _, item := range nestedArr {
			flat = append(flat, Flatten(item)...)
		}
	case interface{}:
		item := nested.(interface{})
		flat = append(flat, item)
	}
	return flat
}
