package main

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}
