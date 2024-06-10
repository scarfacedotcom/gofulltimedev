package main

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type FooStorage struct {
}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct {
}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, store Storage) error {
	return store.Put(id, val)
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	updateValue(1, "petre", s.store)

}
