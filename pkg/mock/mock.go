package mock

import (
	"github.com/brianvoe/gofakeit/v6"
)

type MockOccupation struct {
	Company    string `fake:"{company}" json:"company"`
	Title      string `fake:"{jobtitle}" json:"title"`
	Descriptor string `fake:"{jobdescriptor}" json:"descriptor"`
}

type MockCreditCard struct {
	Type    string `fake:"{creditcardtype}" json:"type"`
	Cvv     int    `fake:"{creditcardcvv}" json:"cvv"`
	Expires string `fake:"{creditcardexp}" json:"expires"`
	Number  string `fake:"{creditcardnumber}" json:"number"`
}

type MockPaymentInfo struct {
	RoutingNumber string         `fake:"{achrouting}" json:"routing_number"`
	AccountNumber string         `fake:"{achaccount}" json:"account_number"`
	SSN           string         `fake:"{ssn}" json:"ssn"`
	CreditCard    MockCreditCard `json:"credit_card"`
}

type MockAddress struct {
	Street    string  `fake:"{streetname}" json:"street"`
	Number    int     `fake:"{streetnumber}" json:"number"`
	City      string  `fake:"{city}" json:"City"`
	State     string  `fake:"{state}" json:"State"`
	Zip       string  `fake:"{zip}" json:"zip"`
	Country   string  `fake:"{country}" json:"country"`
	Latitude  float64 `fake:"{latitude}" json:"latitude"`
	Longitude float64 `fake:"{longitude}" json:"longitude"`
}

type MockUser struct {
	Id             int               `fake:"{number:1111,9999}" json:"id"`
	FirstName      string            `fake:"{firstname}" json:"first_name"`
	LastName       string            `fake:"{lastname}" json:"last_name"`
	Email          string            `fake:"{email}" json:"email"`
	PhoneNumber    string            `fake:"{phoneformatted}" json:"phone_number"`
	PaymentMethods []MockPaymentInfo `fakesize:"2" json:"payment_methods"`
	Occupation     MockOccupation    `json:"occupation"`
	Addresses      []MockAddress     `fakesize:"1" json:"addresses"`
}

func CreateUser() MockUser {
	var user MockUser
	gofakeit.Struct(&user)
	return user
}
