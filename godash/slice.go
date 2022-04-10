package godash

func Map[T any, R any](slice []T, f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any, R any](slice []T, f func(prev R, curr T) R, init R) R {
	result := init
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}
