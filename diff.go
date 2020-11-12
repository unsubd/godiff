package main

func Diff(first []byte, second []byte) map[int]bool {
	indices := make(map[int]bool)
	small := len(first)

	if small > len(second) {
		small = len(second)
	}

	for i := 0; i < small; i++ {
		if first[i] != second[i] {
			indices[i] = true
		}
	}

	return indices
}
