package array_utils

import "errors"

type ZipPair[T, U any] struct {
	X T
	Y U
}

func Zip[T, U any](ts []T, us []U) ([]ZipPair[T, U], error) {

	if len(ts) == 0 || len(us) == 0 {
		return nil, errors.New("cannot zip an empty array")
	}
	zip_length := min(len(ts), len(us))
	var pairs []ZipPair[T, U]
	for i := 0; i < zip_length; i++ {
		pairs = append(pairs, ZipPair[T, U]{ts[i], us[i]})
	}
	return pairs, nil
}
