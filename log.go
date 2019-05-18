package main

import (
	"fmt"
)

// here you can implement log action
// you can send email
// save to database
// send request to log micro service
func sendLog (message string , err error)  {
	fmt.Println(message , err.Error())
}
