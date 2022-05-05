package homework

func sortMapValues(input map[int]string) (result []string) {
	for _, value := range input {
		result = append(result, value)
	}
	return result
}
