package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// check if index exits
func isset(arr []string, index int) bool {
	return (len(arr) > index)
}

func returnWithError(message string) map[string]interface{} {
	return gin.H{
		"message": message,
		"status":  false,
		"data":    "",
	}
}

// set response if micro service not have response
func notHaveResponse() map[string]interface{}   {
	return returnWithError("this micro service not have response body")
}

// get the request headers and pass all to the new request
func  PassAllHeaderToMicroService(req http.Header){
	for key , _ := range req{
		req.Set(key, req.Get(key))
	}
}

func notAuthorized () map[string]interface{} {
	return  returnWithError("you are not Authorized")
}

// check if value in array
func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}