package main

type Putter interface {
	Put(id int, val any) error
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}
type simplePutter struct {
}

func (s *simplePutter) Put(id int, val any) error {
	return nil
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

type Servers struct {
	store Storage
}

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

func main3() {
	s := &Servers{
		store: &FooStorage{},
	}
	updateValue(1, "petre", s.store)

}
