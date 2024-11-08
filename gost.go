package gost

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

type Numeric interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Percent[T1 Numeric, T2 Numeric](num T1, denom T2) float64 {
	return (float64(num) / float64(denom)) * 100.0
}
