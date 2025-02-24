package linkedlist

type List[T comparable] interface {
	Add(values ...T)
	Get(index int) (T, error)
	Remove(index int) error
	Contains(values ...T) bool
	Swap(index1, index2 int) error
	Insert(index int, values ...T) error
	Set(index int, value T) error
}
