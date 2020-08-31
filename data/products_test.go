package data

import "testing"

func TestChecksValidations(t *testing.T) {
	p := &Product{
		Name:  "Kaan",
		Price: 2.34,
		SKU:   "abs-bab-asdf",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
