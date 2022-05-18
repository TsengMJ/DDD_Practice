package opt

import "fmt"

type IOpt interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type Opt struct {
	IOpt IOpt
}

func (o *Opt) GenAndSendOTP(otpLength int) error {
	otp := o.IOpt.genRandomOTP(otpLength)
	o.IOpt.saveOTPCache(otp)
	message := o.IOpt.getMessage(otp)
	err := o.IOpt.sendNotification(message)
	if err != nil {
		return err
	}
	o.IOpt.publishMetric()
	return nil
}

func (o *Opt) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Generating random otp %s\n", randomOTP)
	return randomOTP
}

func (o *Opt) saveOTPCache(otp string) {
	fmt.Printf("Saving otp: %s to cache\n", otp)
}
