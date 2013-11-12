// Copyright 2013 Manish Malik (manishmalik.name)
// All rights reserved.
// Use of this source code is governed by a BSD (3-Clause) License
// that can be found in the LICENSE file.
//
// Currently available APIs:
//
// Initialization:
// 		InitOnePageWithUserPass
//
// Separate Auth APIs:
//		GetNewAuthKey
//

package onepagego

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
)

// LoginResponse represents response of 'login' API
type LoginResponse struct {
	Status int    `json:"status"`
	Message string  `json:"message"`
	Timestamp int  `json:"timestamp"`
	Data LoginRespData `json:"data"`
}

// LoginRespData respresents the 'data' portion of LoginResponse
type LoginRespData struct {
	UID string `json:"uid"`
	Key string `json:"key"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Company string `json:"company"`
	Email string `json:"email"`
	Timezone string `json:"timezone"`
	AccType string `json:"acctype"`
	DateFormat string `json:"dateformat"`
	// TODO:
	// Listing
	// Currency
}

var onepage_app_uid, onepage_auth_key, onepage_api_ver, onepage_version string
var onepage_username, onepage_password string
var onepage_init_done bool

func InitOnePageWithUserPass(api_ver, username, password string) {

	if api_ver != "" && username != "" && password != "" {
		onepage_api_ver = api_ver
		// These variables are reset to blank once they're used
		onepage_username = username
		onepage_password = password

		onepage_init_done = true
	}
}

// GetNewAuthKey function gets a new Auth Key and UID.
// Inputs are (Username string, Password string).
// It returns (UID string, Auth Key string, API status int, Message string).
func GetNewAuthKey(username, password string) (string, string, int, string) {

	jsonobj := new(LoginResponse)

	api_result := callAPI("auth", "login="+username+"&password="+password)

	jsonerr := json.Unmarshal([]byte(api_result), jsonobj)

	if jsonerr != nil {
		jsonobj.Message = "Invalid JSON: " + jsonerr.Error()
	}

	return jsonobj.Data.UID, jsonobj.Data.Key, jsonobj.Status, jsonobj.Message
}

// callAPI Internal function handling the REST API.
// This uses the .json format for all calls.
func callAPI(apicall, apidata string) string {

	// Check if we have auth token available.
	// If not, let's first authenticate and retrieve it.
	if apicall != "auth" && onepage_auth_key == "" {

		new_auth_uid, new_auth_key, new_auth_status, _ := GetNewAuthKey(onepage_username, onepage_password)

		onepage_username = ""
		onepage_password = ""

		if new_auth_key == "" || new_auth_status != 0 {
			api_result := "{\"success\":false, \"message\":\"Unable to get a valid Auth Key from API.\" }"
			return api_result
		} else {
			onepage_app_uid = new_auth_uid
			onepage_auth_key = new_auth_key
		}
	}

	api_result := ""
	client := &http.Client{}
	api_method := "GET" // overridden later
	var param_data []byte
	var param_reader *bytes.Reader
	param_data = ([]byte)("")

	// Decide on the HTTP Method and Parameters
	if apicall == "auth" {
		api_method = "POST"
		apicall = "auth/login.json"
		param_data = ([]byte)(apidata)
	} else {
		api_method = "GET"
	}

	// Make the API URL to call
	api_url := "https://app.onepagecrm.com/api/" + apicall

	param_reader = bytes.NewReader(param_data)

	req, err := http.NewRequest(api_method, api_url, param_reader)
	if err == nil {

		if apicall != "auth" {
			ts := strconv.FormatInt(time.Now().Unix(), 10)

			req.Header.Add("X-OnePageCRM-UID", onepage_app_uid)
			req.Header.Add("X-OnePageCRM-TS", ts)
			req.Header.Add("X-OnePageCRM-Auth", "") // TODO: Calculate Request Signature
		}

		resp, resperr := client.Do(req)
		if resperr != nil {
			api_result = "{\"success\":false, \"message\":\"Error connecting to or retrieving response from API URL. Please check connectivity. API URL: " + api_url + "\" }"
		} else {
			defer resp.Body.Close()
			bodybytes, _ := ioutil.ReadAll(resp.Body)
			api_result = string(bodybytes)
		}
	}

	return api_result
}
