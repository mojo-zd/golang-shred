package observer

import "testing"

func TestObserver(t *testing.T) {
	subMsg := NewSubjectMsg().(*subjectMsg)
	observers := []*Observer{
		{UUID: "1"},
		{UUID: "2"},
		{UUID: "3"},
		{UUID: "2"},
	}
	subMsg.Registry(observers...)
	subMsg.SetContext("msg context add")
	subMsg.Notify()
	t.Log("===============")
	subMsg.SetContext("msg context update")
	subMsg.Notify()
	t.Log("===============")
	subMsg.SetContext("msg context delete")
	subMsg.Notify()

	t.Log("===============")
	mail := NewSubjectMail().(*subjectMail)
	mail.Registry(observers...)
	mail.SetContext("mail context add")
	mail.Notify()

	t.Log("===============")
	mail.SetContext("mail context update")
	mail.Notify()

	t.Log("===============")
	mail.SetContext("mail context delete")
	mail.Notify()
}
