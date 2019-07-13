package system

import (
	uos "github.com/uoscanada/uos-go"
)

func NewBuyRAM(payer, receiver uos.AccountName, uosQuantity uint64) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("buyram"),
		Authorization: []uos.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: uos.NewUOSAsset(int64(uosQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `uosio.system::buyram` action.
type BuyRAM struct {
	Payer    uos.AccountName `json:"payer"`
	Receiver uos.AccountName `json:"receiver"`
	Quantity uos.Asset       `json:"quant"` // specified in UOS
}