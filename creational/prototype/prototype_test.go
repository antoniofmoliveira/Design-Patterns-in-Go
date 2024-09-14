package prototype

import "testing"

// This is a Go test function named `TestClone` that tests the cloning functionality of a shirt cache. Here's a succinct explanation:
// 1. It retrieves a shirt cache using `GetShirtsCloner()` and checks if it's not nil.
// 2. It clones a white shirt from the cache using `GetClone(White)` and checks if the cloned shirt is not equal to the original prototype.
// 3. It modifies the cloned shirt's SKU and then clones another white shirt from the cache.
// 4. It checks if the two cloned shirts have different SKUs and are not the same object.
// 5. It logs information about the cloned shirts, including their memory positions.
// 6. It tests cloning other shirt models (Black, Blue) and an invalid model (10), expecting an error for the invalid model.
// The test ensures that the shirt cache correctly clones shirts, modifies them independently, and handles invalid models.
func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to Shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: The memory positions of the shirts are different %p != %p \n\n", &shirt1, &shirt2)

	_, err = shirtCache.GetClone(Black)
	if err != nil {
		t.Fatal(err)
	}
	_, err = shirtCache.GetClone(Blue)
	if err != nil {
		t.Fatal(err)
	}
	_, err = shirtCache.GetClone(10)
	if err == nil {
		t.Fatal("An error must be returned for an invalid model")
	}
}
