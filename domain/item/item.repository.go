package item

type ItemRepository interface {
	Get() ([]Item, error)
}
