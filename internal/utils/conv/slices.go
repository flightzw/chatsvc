package conv

type SliceMap[T any, K string | int | int32, V any] []T

func (s SliceMap[T, K, V]) Map(f func(T) [2]any) map[K]V {
	res := map[K]V{}
	for _, item := range s {
		keyval := f(item)
		res[keyval[0].(K)] = keyval[1].(V)
	}
	return res
}
