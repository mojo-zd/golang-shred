package templatemethod

import "fmt"

type BankTemplate interface {
	TakeQueue() int     //排队
	Transaction() error //办理业务
	Feedback() string   // 反馈
	Process() error     //执行业务
}

type BankTemplateBasic struct {
	BankTemplate
}

// 排队实现
func (btb *BankTemplateBasic) TakeQueue() int {
	num := 1
	fmt.Printf("取到排队编号%d\n", num)
	return num
}

// 反馈
func (btb *BankTemplateBasic) Feedback() string {
	fmt.Printf("接收反馈信息...\n")
	return ""
}

func (btb *BankTemplateBasic) Process() error {
	btb.TakeQueue()
	err := btb.Transaction()
	btb.Feedback()
	return err
}

type nongYeBankTemplateImpl struct {
	BankTemplateBasic
}

func (bti *nongYeBankTemplateImpl) Transaction() error {
	fmt.Println("农业银行模板业务实现...")
	return nil
}

func NewNongYeBankTemplate() BankTemplate {
	impl := &nongYeBankTemplateImpl{}
	impl.BankTemplateBasic = BankTemplateBasic{impl}
	return impl
}

type gongShangBankTemplateImpl struct {
	BankTemplateBasic
}

func (*gongShangBankTemplateImpl) Transaction() error {
	fmt.Println("工商银行业务实现...")
	return nil
}

func NewGongShangBankTemplateImpl() BankTemplate {
	bank := &gongShangBankTemplateImpl{}
	bank.BankTemplateBasic = BankTemplateBasic{bank}
	return bank
}
