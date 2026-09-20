package realty

import "errors"

type Property struct {
	ListingID string
	Address   string
	Price     float64
	Bedrooms  int
	IsSold    bool
}

var listings = make(map[string]Property)

func ListProperty(p Property) error {
	if p.ListingID == "" {
		return errors.New("listing ID cannot be empty")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if p.Bedrooms < 1 {
		return errors.New("bedrooms must be at least 1")
	}
	if _, exists := listings[p.ListingID]; exists {
		return errors.New("listing already exists: " + p.ListingID)
	}
	listings[p.ListingID] = p
	return nil
}
func GetProperty(id string) (*Property, error) {
	p, exists := listings[id]
	if !exists {
		return nil, errors.New("property not found: " + id)
	}
	return &p, nil
}
func UpdateListingPrice(id string, newPrice float64) error {
	p, exists := listings[id]
	if !exists {
		return errors.New("property not found: " + id)
	}
	if newPrice <= 0 {
		return errors.New("price must be greater than 0")
	}
	p.Price = newPrice
	listings[id] = p
	return nil
}
func DelistProperty(id string) error {
	if _, exists := listings[id]; !exists {
		return errors.New("property not found: " + id)
	}
	delete(listings, id)
	return nil
}
