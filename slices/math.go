package slices

const testVersion = 1

// AddUint8s adds each element of first to the respective element of second and returns a new slice holding the result.
// If the length of the inputs are different, the result will be the length of the shorter input.
// The resulting indexes
func AddUint8s(first, second []uint8) []uint8 {
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

// AddToUint8Slice adds toAdd to each element of nums
func AddToUint8Slice(nums []uint8, toAdd uint8) {
	for i := range nums {
		nums[i] += toAdd
	}
}

// SubtractFromUint8Slice subtracts toSubtract from each element of nums
func SubtractFromUint8Slice(nums []uint8, toSubtract uint8) {
	for i := range nums {
		nums[i] -= toSubtract
	}
}

// MultiplyForUint8Slice multiplies each element of nums by multiplier
func MultiplyForUint8Slice(nums []uint8, multiplier uint8) {
	for i := range nums {
		nums[i] *= multiplier
	}
}
