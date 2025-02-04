package util

func GetOptional[T any](inp []T) T {
	if len(inp) == 0 {
		return *new(T)
	}
	return inp[0]
}

func Coalesce[T comparable](inp []T) T {
	defaultVal := *new(T)
	return CoalesceWithDefault(defaultVal, inp)
}

func CoalesceWithDefault[T comparable](defaultVal T, inp []T) T {
	for _, i := range inp {
		if i != defaultVal {
			return i
		}
	}
	return defaultVal
}
