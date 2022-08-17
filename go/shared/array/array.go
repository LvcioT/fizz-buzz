package array

import "errors"

func Map[T_Source any, T_Dest any](data []T_Source, f func(T_Source) T_Dest) []T_Dest {
	mapped := make([]T_Dest, len(data))

	for i, item := range data {
		mapped[i] = f(item)
	}

	return mapped
}

func Each[T any](data []T, f func(int, T)) {
	for i, item := range data {
		f(i, item)
	}
}

func Filter[T any](data []T, f func(T) bool) []T {
	filtered := make([]T, 0)

	for _, item := range data {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func First[T any](data []T, f func(T) bool) (T, error) {
	var filtered T

	for _, item := range data {
		if f(item) {
			return item, nil
		}
	}

	return filtered, errors.New("item not found")
}

func In[T comparable](data []T, value T) bool {
	for _, item := range data {
		if item == value {
			return true
		}

	}

	return false
}
