package main

import (
	"fmt"
	"template_method/opt"
)

func main() {
	smsOTP := &opt.SMS{}
	o := opt.Opt{
		IOpt: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &opt.Email{}
	o = opt.Opt{
		IOpt: emailOTP,
	}
	o.GenAndSendOTP(4)

}
