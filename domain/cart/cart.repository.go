package cart

type CartRepository interface {
	Add(int, int, int) error
	Get(int) ([]Cart, error)
}
