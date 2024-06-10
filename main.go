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

type Server struct {
	store Storage
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	s.store.Get(1)
	s.store.Put(1, "foo")

}
