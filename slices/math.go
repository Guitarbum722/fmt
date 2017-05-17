package slices

const testVersion = 1

// AddUint8s adds each element of first to the respective element of second and returns a new slice holding the result.
// If the length of the inputs are different, the result will be the length of the shorter input.
// The resulting indexes
func AddUint8s(first []uint8, second []uint8) []uint8 {
	var result []uint8

	switch len(first) <= len(second) {
	case true:
		for i := range first {
			result = append(result, first[i]+second[i])
		}
	case false:
		for i := range second {
			result = append(result, first[i]+second[i])
		}
	}

	return result
}
