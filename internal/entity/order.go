package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func CreateOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) Validate() error {
	if order.ID == "" {
		return errors.New("ID is invalid")
	}

	if order.Price <= 0 {
		return errors.New("price is invalid")
	}

	if order.Tax <= 0 {
		return errors.New("tax is invalid")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() error {
	order.FinalPrice = order.Price + order.Tax

	err := order.Validate()

	if err != nil {
		return err
	}

	return nil
}
