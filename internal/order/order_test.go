package order

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestComputeTotal(t *testing.T) {
	o:=Order{
		ID: "45",
		CurrencyAlphaCode: "EUR",
		Items: []Item{
			{
				ID: "458",
				Quantity: 2,
				UnitPrice: money.New(100,"EUR"),
			},
		},
	}
	total,err:=o.ComputeTotal()
	assert.NoError(t,err)
	assert.Equal(t,int64(200), total.Amount())
	assert.Equal(t, "EUR", total.Currency().Code)
}