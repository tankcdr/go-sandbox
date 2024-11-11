package hash

// *db2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func Hash_djb2(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}
