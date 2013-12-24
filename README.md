# go-gmail

go-gmail is an up-and-coming GMail library in [Golang](http://golang.org/). For now it only supports sending e-mail using authenticated SMTP.


## Installation

Assuming you have a working Go environment, installation is simple:

    go get github.com/njern/gogmail

You can take a look at the documentation locally with:

	godoc github.com/njern/gogmail

The included tests in `gogmail_test.go` also illustrate usage of the package. 

**Note:** You must enter a valid GMail address / password in `gogmail_test.go` or the tests will fail! 
For obvious reasons I did not leave my own address / password in the test file... :)

## Usage
    import "github.com/njern/gogmail"
    import "log"
    
    gmail := gogmail.GmailConnection("gmailuser@gmail.com", "my_password")

    err := gmail.SendMail([]string{"niclas@walkbase.com"}, "Hello go-gmail!", "Looks like go-gmail is working just fine!")
	if err != nil {
		log.Fatal("Sending a GMail e-mail failed with error: ", err)
	}
    
## Future plans

* Support for retrieving e-mail, folders, etc etc

## How can you help?

* Let me know if you're using go-gmail by dropping me a line at [github user name] at walkbase.com
* Let me know about any bugs / annoyances the same way