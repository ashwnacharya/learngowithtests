package generic_arrays_slices

type Transaction struct {
	From string
	To string
	Sum float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func BalanceFor(transactions []Transaction, name string) (balance float64) {
	
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		} else if t.To == name {
			return currentBalance + t.Sum
		} else {
			return currentBalance
		}
	}

	return Reduce(transactions, adjustBalance, 0.0)
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

