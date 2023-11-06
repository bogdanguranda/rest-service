package db

type DB interface {
	GetInput() ([]int, error)
}
