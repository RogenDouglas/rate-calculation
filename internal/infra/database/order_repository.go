package database

import (
	"database/sql"

	"github.com/rate-calculation/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func CreateOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO ORDERS (ID, PRICE, TAX, FINAL_PRICE) VALUES (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM ORDERS").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
