package design_pattern

import (
	"fmt"
	"strconv"
	"time"
)

type PaymentContext struct {
	Name    string
	CardID  string
	Money   int
	payment PaymentStrategy
}

func NewPaymentContext(name, cardID string, money int, payment PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		Name:    name,
		CardID:  cardID,
		Money:   money,
		payment: payment,
	}
}

func (context *PaymentContext) Pay() {
	context.payment.Pay(context)
}

type PaymentStrategy interface {
	Pay(context *PaymentContext)
	String(context *PaymentContext) string
}

type Cash struct {
}

func (cash *Cash) Pay(context *PaymentContext) {
	fmt.Println(cash.String(context))
}

func (cash *Cash) String(context *PaymentContext) string {
	return "Name:" + context.Name +
		",CardID" + context.CardID +
		",Money:" + strconv.Itoa(context.Money) +
		",pay method: cash" +
		",time:" +time.Now().String()
}

type Bank struct {
}

func (bank *Bank) Pay(context *PaymentContext) {
	fmt.Println(bank.String(context))
}

func (bank *Bank) String(context *PaymentContext) string {
	return "Name:" + context.Name +
		",CardID" + context.CardID +
		",Money:" + strconv.Itoa(context.Money) +
		",pay method: bank" +
		",time:" +time.Now().String()
}
