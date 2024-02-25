package service

import (
	"midtrans-example/model"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

const (
	MERCHANT_ID = "G568921549"
	CLIENT_KEY  = "SB-Mid-client-l9X94jCTxlUnVGB5"
	SERVER_KEY  = "SB-Mid-server-VbnO3JrBpFqYeImbcGZw18o8"
)

func Create(request model.MidtransRequest) *model.MidtransResponse {
	var snapClient snap.Client

	snapClient.New(SERVER_KEY, midtrans.Sandbox)

	user_id := strconv.Itoa(request.UserID)

	custAddress := &midtrans.CustomerAddress{
		FName:       "foo",
		LName:       "bar",
		Phone:       "0987654321",
		Address:     "Coral Street",
		City:        "Bikini Bottom",
		Postcode:    "69969",
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + user_id + "-",
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "Foo",
			LName:    "Bar",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property" + request.ItemID,
				Qty:   1,
				Price: request.Amount,
				Name:  request.ItemName,
			},
		},
	}

	resp, err := snapClient.CreateTransaction(req)

	if err != nil {
		return nil
	}

	return &model.MidtransResponse{
		Token:       resp.Token,
		RedirectURL: resp.RedirectURL,
	}
}
