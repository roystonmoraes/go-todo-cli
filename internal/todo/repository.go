package todo

type Repository interface {
	Load() ([]Todo, error)
	Save([]Todo) error
}
