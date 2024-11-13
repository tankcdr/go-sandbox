package hash

const EMPTY = ""

const DELETED = true
const NOT_DELETED = !DELETED
const USED = NOT_DELETED

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

// Jenkins one_at_a_time hash function.
// See https://en.wikipedia.org/wiki/Jenkins_hash_function
func Hash_jenkins(value string) int {
	hash := 0
	for _, ch := range value {
		hash += int(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}

	// Make sure the result is not 0.
	if hash == 0 {
		hash = 1
	}
	return hash
}
