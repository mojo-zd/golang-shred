package templatemethod

import "testing"

func TestBankTemplate(t *testing.T) {
	nybank := NewNongYeBankTemplate()
	nybank.Process()
	t.Log("=========")
	gsbank := NewGongShangBankTemplateImpl()
	gsbank.Process()
}
