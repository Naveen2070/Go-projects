package expenseModel

type Expense struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}

var Expenses []Expense
