package prototype

import (
	"errors"
	"fmt"
)

// `ShirtCloner` is an interface that defines a contract for cloning shirts.
// **Method:**
// * `GetClone(s int) (ItemInfoGetter, error)`: This method takes an integer `s` as input and returns a cloned shirt as an `ItemInfoGetter` object, along with an error if the cloning operation fails.
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

// * It defines an interface named `ItemInfoGetter` that has one method:
//   - `GetInfo() string`: This method returns a string containing information about an item. Any type that implements this interface must provide an implementation for this method.
type ItemInfoGetter interface {
	GetInfo() string
}

// `ShirtsCache` is a struct that implements the `ShirtCloner` interface.
type ShirtsCache struct{}

// `ShirtColor` is an alias for a byte type.
type ShirtColor byte

// This struct definition creates a `Shirt` struct with three fields: `Price` (a float32 representing the price of the shirt), `SKU` (a string representing the stock keeping unit of the shirt), and `Color` (a custom type `ShirtColor` representing the color of the shirt). 
// The `Shirt` struct does not have any methods defined in this code snippet. It is a simple data structure used to store information about a shirt.
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// `ShirtColor` is a custom type that represents the color of a shirt. It has three possible values: `White`, `Black`, and `Blue`.
const (
	White = 1
	Black = 2
	Blue  = 3
)

// simulates the database
var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}
var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// GetShirtsCloner returns a new instance of a `ShirtCloner`, which is a struct that implements the `ShirtCloner` interface. It is used to clone prototype shirts.
func GetShirtsCloner() ShirtCloner {
	shirtsCache := new(ShirtsCache)
	return shirtsCache
}

// GetClone returns a new instance of the requested shirt, or an error if the requested model is not recognized.
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

// GetInfo returns a string representation of the shirt, including its SKU, color, and price.
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.GetPrice())
}

// GetPrice returns the price of the shirt.
func (i *Shirt) GetPrice() float32 {
	return i.Price
}
