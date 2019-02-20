package signedpay_api

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

func TestApi_Charge(t *testing.T) {
	api := New("unicorn", "20c20ee3-4173-4daa-87e5-dbcce8c7949d", true)
	//orderID:= uuid.New()
	charge, err := api.Charge(&ChargeRequest{
		Amount:           1000,
		Currency:         "USD",
		CustomerEmail:    "test@test.com",
		IPAddress:        "8.8.8.8",
		Platform:         "WEB",
		OrderID:          "11",
		OrderDescription: "test description",
		CardCvv:          "967",
		CardExpMonth:     "03",
		CardExpYear:      "2021",
		CardNumber:       "4532456618142692",
	})

	if err != nil {
		if e, ok := err.(PayError); ok {
			logrus.Error(e)
		} else {
			logrus.Fatal(err)
		}
		return
	}
	logrus.Info(charge.Transactions)
}

func TestApi_Recurring(t *testing.T) {
	api := New("unicorn", "20c20ee3-4173-4daa-87e5-dbcce8c7949d", true)
	api.Recurring(&RecurringRequest{
		Amount:           1000,
		Currency:         "USD",
		CustomerEmail:    "test@test.com",
		IPAddress:        "8.8.8.8",
		Platform:         "WEB",
		OrderID:          "10",
		OrderDescription: "test description",
		RecurringToken:   "ecb00e3333ca3e9eff39e9e87eee3d5a8b2733a4a834a6cc0fa6fbb6e0dee260fb919e30ea0ca9bece1f2eb8b00c6d46143a",
	})
}

func TestApi_Refund(t *testing.T) {
	api := New("unicorn", "20c20ee3-4173-4daa-87e5-dbcce8c7949d", true)
	response, err := api.Refund(&RefundRequest{OrderID: "2", Amount: 500})
	if err != nil {
		if e, ok := err.(PayError); ok {
			logrus.Error(e)
		} else {
			logrus.Fatal(err)
		}
		return
	}
	logrus.Info(response.Order)

}

func TestApi_CheckOrderStatus(t *testing.T) {
	api := New("unicorn", "20c20ee3-4173-4daa-87e5-dbcce8c7949d", true)
	response, err := api.Status(&StatusRequest{OrderID: "2"})
	if err != nil {
		if e, ok := err.(PayError); ok {
			logrus.Error(e)
		} else {
			logrus.Fatal(err)
		}
		return
	}
	logrus.Info(response.Transactions)
}
