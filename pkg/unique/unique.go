package unique

func Unique[T comparable](s []T) (r []T) {
	var unique = map[T]struct{}{}

	for _, v := range s {
		if _, exists := unique[v]; !exists {
			unique[v] = struct{}{}
			r = append(r, v)
		}
	}

	return r
}
