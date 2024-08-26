package generics

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[key comparable, value int64 | float64](m map[key]value) value {
	var s value

	for _, v := range m {
		s += v
	}
	return s
}
