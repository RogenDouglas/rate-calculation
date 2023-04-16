package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDIsValid(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "ID is invalid")

}

func TestPriceIsValid(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.Validate(), "price is invalid")

}

func TestTaxIsValid(t *testing.T) {
	order := Order{ID: "1", Price: 3.0}
	assert.Error(t, order.Validate(), "tax is invalid")

}

func TestWithAllalidParams(t *testing.T) {
	order := Order{ID: "1", Price: 3.0, Tax: 2.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 3.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 5.0, order.FinalPrice)

}
