package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/L0/internal/app/model"
)

func TestOrder_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		o       func() *model.Order
		isValid bool
	}{
		{
			name: "valid",
			o: func() *model.Order {
				return model.TestOrder(t)
			},
			isValid: true,
		},
		{
			name: "empty_id",
			o: func() *model.Order {
				o := model.TestOrder(t)
				o.OrderId = ""
				return o
			},
			isValid: false,
		},
		{
			name: "locale_notCorrect",
			o: func() *model.Order {
				o := model.TestOrder(t)
				o.Locale = "russkia"
				return o
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.o().Validate())
			} else {
				assert.Error(t, tc.o().Validate())
			}
		})
	}
}
