package beerscli

// Store representation of beer into data struct
type Store struct {
	StoreID int    `json:"product_id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// StoreRepo definiton of methods to access stores data
type StoreRepo interface {
	GetStores() ([]Store, error)
}

// NewStore initialize struct Store
func NewStore(storeID int, name, country string) (b Store) {
	b = Store{
		StoreID: storeID,
		Name:    name,
		Country: country,
	}
	return
}
