package commonutils

func NextCircularIndex(arrLen int, currentPosition int, step int) int {
	nextIdx := currentPosition + step
	if nextIdx >= arrLen {
		nextIdx %= arrLen
	}
	return nextIdx
}