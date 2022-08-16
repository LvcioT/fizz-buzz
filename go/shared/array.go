package shared

import "errors"

func Array_map[i_item any, r_item any](data []i_item, f func(i_item) r_item) []r_item {
	mapped := make([]r_item, len(data))

	for i, item := range data {
		mapped[i] = f(item)
	}

	return mapped
}

func Array_each[item any](data []item, f func(int, item)) {
	for i, item := range data {
		f(i, item)
	}
}

func Array_filter[i_item any](data []i_item, f func(i_item) bool) []i_item {
	filtered := make([]i_item, 0)

	for _, item := range data {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func Array_first[i_item any](data []i_item, f func(i_item) bool) (i_item, error) {
	var filtered i_item

	for _, item := range data {
		if f(item) {
			return item, nil
		}
	}

	return filtered, errors.New("item not found")
}

func Array_in[T comparable](data []T, value T) bool {
	for _, item := range data {
		if item == value {
			return true
		}

	}

	return false
}
