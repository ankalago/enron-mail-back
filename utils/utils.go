package utils

func RemoveNullValue(slice []interface{}) []interface{} {
	var output []interface{}
	for _, element := range slice {
		if element != nil { //if condition satisfies add the elements in new slice
			output = append(output, element)
		}
	}
	return output //slice with no nil-values
}
