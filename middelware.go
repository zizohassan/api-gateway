package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// here we will handel auth routes in this micro service
// if you want to add more go to routes file
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, notAuthorized())
			c.Abort()
			return
		}
		/// implement your authorization action here
		fmt.Println("here is auth : " + authorization)
		c.Next()
	}
}

// if request come here it must not one of this micro service routing
// we check the micro service name with the first request uri segment
// then check the action by check the second segment of request uri
// if action define in action slice the get way will automatic apply
// auth middleware if not it will send the request with header to the
// define micro service and will get response and send it back to user
func handelOtherMicroServiceRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		// split to get micro service name
		// return 404 if we not found segment in the uri
		var segments []string
		url := c.Request.RequestURI
		// check if url has query params split
		// else just split only "/"
		if strings.Contains(url , "?"){
			segments = strings.Split(url, "?")
			segments = strings.Split(segments[0], "/")
		}else{
			segments = strings.Split(url, "/")
		}
		if !isset(segments, 1) {
			c.JSON(http.StatusNotFound, returnWithError("wrong url"))
			c.Abort()
			return
		}
		//check authenticated actions
		if isset(segments, 2) {
			auth := checkIfActionNeedAuth(segments[2], segments[1])
			if auth {
				authorization := c.GetHeader("Authorization")
				if authorization == "" {
					c.JSON(http.StatusUnauthorized, notAuthorized())
					c.Abort()
					return
				}
				/// implement your authorization action here
				fmt.Println(authorization)
			}
		}
		// check if the segment have value in our
		// micro services url map
		microServiceUrl, find := getMicroServiceUrl(segments[1])
		if find == false {
			c.JSON(http.StatusNotFound, returnWithError("we not found micro service to serve"))
			c.Abort()
			return
		}
		// check the request method
		// then use the match function with method
		method := strings.ToLower(c.Request.Method)
		url = microServiceUrl + url
		switch method {
		case "post":
			post(c, url)
			break
		case "get":
			get(c, url)
			break
		}
	}
}

// check if the segment of url have the same index key with
// micro service then return the url of micro service
func getMicroServiceUrl(firstSegment string) (string, bool) {
	urls := microServiceMap()
	if _, ok := urls[firstSegment]; ok {
		return urls[firstSegment], true
	}
	return "", false
}

// these action if where in the url
// you must put the token Authorization to check if you
// are auth or not
func checkIfActionNeedAuth(action string, microService string) bool {
	var actions = authActionsSlice()
	if _, ok := actions[microService]; ok {
		exists, _ :=in_array(action , actions[microService])
		if exists {
			return true
		}
	}
	return false
}

// cross origin
// allow only get post requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Next()
	}
}
