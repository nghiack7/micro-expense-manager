package models

import "time"

type Expense struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
