package repositories

import (
	"github.com/nghiack7/micro-expense-manager/expense-service/pkg/models"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository { return &expenseRepository{db: db} }

func (e *expenseRepository) CreateNewExpense(expense *models.Expense) (*models.Expense, error) {
	err := e.db.Create(expense).Error
	if err != nil {
		return nil, err
	}
	return expense, nil
}
