package homework

func sortMapValues(input map[int]string) (result []string) {
	for key := range input {
		result = append(result, input[key])
	}
	return result
}
