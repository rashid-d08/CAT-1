package realty

import (
	"encoding/json"
	"errors"
	"os"
)

type Property struct {
	ListingID string  `json:"listing_id"`
	Address   string  `json:"address"`
	Price     float64 `json:"price"`
	Bedrooms  int     `json:"bedrooms"`
	IsSold    bool    `json:"is_sold"`
}

const databaseFile = "database/database.json"

var listings = make(map[string]Property)

func init() {
	loadDatabase()
}

func loadDatabase() {

	data, err := os.ReadFile(databaseFile)

	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}

	json.Unmarshal(data, &listings)
}

func saveDatabase() error {

	data, err := json.MarshalIndent(listings, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(databaseFile, data, 0644)
}

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

	return saveDatabase()
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

	return saveDatabase()
}

func DelistProperty(id string) error {

	if _, exists := listings[id]; !exists {
		return errors.New("property not found: " + id)
	}

	delete(listings, id)

	return saveDatabase()
}
