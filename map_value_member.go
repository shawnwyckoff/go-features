package main

import "fmt"

type SubAccount struct {
	Amount       float64 `json:"Amount"` // 总数，包括了FrozenAmount
	FrozenAmount float64 `json:"Locked"`
}

type Account struct {
	SubAccounts map[string]SubAccount `json:"Spot"`
}

// can't change member of map NON-Pointer value like this
// acc.Spot["USD"].Amount = 1000000
// but you can assign new map value
// also, you can change member of value of map[string]*SubAccount

func main() {
	acc := Account{SubAccounts: make(map[string]SubAccount)}

	acc.SubAccounts["USD"] = SubAccount{Amount: 1000000, FrozenAmount: 0}
	fmt.Println(acc)
}
