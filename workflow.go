package app

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// TaskQueueName adalah antrian tugas yang digunakan oleh worker
const TaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

// TransferInput adalah data input untuk workflow
type TransferInput struct {
	FromAccount string
	ToAccount   string
	Amount      float64
	ReferenceID string
}

// TransferWorkflow adalah logika orkestrasi (Workflow)
func TransferWorkflow(ctx workflow.Context, input TransferInput) (string, error) {
	// 1. Tentukan opsi untuk Activity (misal: timeout)
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	logger := workflow.GetLogger(ctx)
	logger.Info("TransferWorkflow started", "ReferenceID", input.ReferenceID)

	// 2. Jalankan Activity pertama (Tarik Saldo)
	var withdrawResult string
	err := workflow.ExecuteActivity(ctx, WithdrawActivity, input).Get(ctx, &withdrawResult)
	if err != nil {
		return "", err
	}

	// 3. Jalankan Activity kedua (Setor Saldo)
	var depositResult string
	err = workflow.ExecuteActivity(ctx, DepositActivity, input).Get(ctx, &depositResult)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Transfer %f from %s to %s completed successfully (Ref: %s)", 
		input.Amount, input.FromAccount, input.ToAccount, input.ReferenceID)
	
	return result, nil
}

// WithdrawActivity adalah langkah menarik saldo (Activity)
func WithdrawActivity(ctx context.Context, input TransferInput) (string, error) {
	fmt.Printf("Activity: Withdrawing %f from account %s...\n", input.Amount, input.FromAccount)
	// Simulasi proses database
	time.Sleep(1 * time.Second)
	return "Withdraw success", nil
}

// DepositActivity adalah langkah setor saldo (Activity)
func DepositActivity(ctx context.Context, input TransferInput) (string, error) {
	fmt.Printf("Activity: Depositing %f to account %s...\n", input.Amount, input.ToAccount)
	// Simulasi proses database
	time.Sleep(1 * time.Second)
	return "Deposit success", nil
}
