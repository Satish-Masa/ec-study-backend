package item

type ItemRepository interface {
	Get() ([]Item, error)
	Find(int) (Item, error)
}
