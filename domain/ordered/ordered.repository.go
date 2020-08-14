package ordered

type OrderedRepository interface {
	Add(int, int, int) error
}
