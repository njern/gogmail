package gogmail

import (
	"testing"
)

const (
	GMAIL_ADDRESS  = "your_email_address@gmail.com"
	GMAIL_PASSWORD = "your_password"
)

func TestSendingGmail(t *testing.T) {
	gmail := GmailConnection(GMAIL_ADDRESS, GMAIL_PASSWORD)

	err := gmail.SendMail([]string{"niclas@walkbase.com"}, "Hello go-gmail!", "Looks like the unit test is running just fine.")
	if err != nil {
		t.Error("Sending a gmail e-mail failed with error: ", err)
	}
}

func TestSendingGmailWithInvalidUser(t *testing.T) {
	gmail := GmailConnection("complete", "bullshit")

	err := gmail.SendMail([]string{"niclas@walkbase.com"}, "Hello go-gmail!", "This really should not be working.")
	if err == nil {
		t.Error("Somehow we managed to send an e-mail with invalid logins: ", err)
	}
}
