package stats

import (
	"github.com/KomGitHub/bank/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	amount := types.Money(0)
	for _, payment := range payments {
		amount += payment.Amount
	}
	return amount / types.Money(len(payments))
}