package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * take request url
 * add all header send to the new request
 * send request to micro service
 * check if response return empty
 * return with micro service response
 */

func get(g *gin.Context, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		sendLog("Error reading request. ", err)
	}
	// get requested header
	// add all send header to send request
	PassAllHeaderToMicroService(g.Request.Header)
	// set time out
	client := &http.Client{Timeout: time.Second * 10}
	// send request
	response, err := client.Do(req)
	// check if there any drop connections
	if err != nil {
		sendLog("Error reading response. ", err)
	}
	defer response.Body.Close()
	// convert result to json again
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	// some times micro service may be not have body to parse
	// so we handel this
	if len(result) == 0 {
		g.JSON(response.StatusCode, notHaveResponse())
		return
	}
	g.JSON(response.StatusCode, result)
	return
}

/**
 * take request data
 * take request url
 * add all header send to the new request
 * send request to micro service
 * check if response return empty
 * return with micro service response
 */

func post(g *gin.Context, url string) {
	//get request data first
	body, err := g.GetRawData()
	if err != nil {
		sendLog("Body parse Error " ,err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		sendLog("Error reading request. ", err)
	}
	// get requested header
	// add all send header to send request
	PassAllHeaderToMicroService(g.Request.Header)
	// set time out
	client := &http.Client{Timeout: time.Second * 10}
	// send request
	response, err := client.Do(req)
	// check if there any drop connections
	if err != nil {
		sendLog("Error reading response. ", err)
	}
	defer response.Body.Close()
	// convert result to json again
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	// some times micro service may be not have body to parse
	// so we handel this
	if len(result) == 0 {
		g.JSON(response.StatusCode, notHaveResponse())
		return
	}
	g.JSON(response.StatusCode, result)
	return
}
