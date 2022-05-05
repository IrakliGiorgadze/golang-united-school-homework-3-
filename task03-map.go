package homework

// func sortMapValues(input map[int]string) (result []string) {
// 	for _, value := range input {
// 		result = append(result, value)
// 	}
// 	return result
// }
import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//return map values in sorted order using sort
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		result = append(result, input[k])
	}
	return result

}
