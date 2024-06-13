package main

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func main6() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("scar", 1)
	m1.Insert("face", 2)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.99)
	m2.Insert(2, 100.889)
}
