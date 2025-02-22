package service

import (
	"context"
	"fmt"

	expensepb "github.com/Naveen2070/Go-projects/go-grpc/common/api"
)

type ExpenseServiceServer struct {
	expensepb.UnimplementedExpenseServiceServer
	expenses []*expensepb.Expense
}

// HealthServiceServer is the implementation of the HealthService
type HealthServiceServer struct {
	expensepb.UnimplementedHealthServiceServer
}

// Check implements the HealthService Check method
func (h *HealthServiceServer) Check(ctx context.Context, req *expensepb.HealthCheckRequest) (*expensepb.HealthCheckResponse, error) {
	return &expensepb.HealthCheckResponse{Status: "SERVING"}, nil
}

func (s *ExpenseServiceServer) CreateExpense(ctx context.Context, req *expensepb.CreateExpenseRequest) (*expensepb.CreateExpenseResponse, error) {
	expense := expensepb.Expense{
		Id:       fmt.Sprintf("%d", len(s.expenses)+1),
		Title:    req.Title,
		Amount:   req.Amount,
		Category: req.Category,
	}
	s.expenses = append(s.expenses, &expense)
	return &expensepb.CreateExpenseResponse{
		Id:      expense.Id,
		Message: "Expense added successfully",
	}, nil
}

func (s *ExpenseServiceServer) GetExpenses(ctx context.Context, req *expensepb.GetExpensesRequest) (*expensepb.GetExpensesResponse, error) {
	return &expensepb.GetExpensesResponse{Expenses: s.expenses}, nil
}
