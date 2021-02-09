package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
)


const CONSUMER_KEY = "consumerkey"
const CONSUMER_SECRET = "consumser_secret"
const API_URL = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

func main(){

}

func getMPESAToken() (interface{}, error) {

	//response, err := cl
	client := &http.Client{
		Jar:           http.CookieJar(),
		CheckRedirect: redirectPolicyFunc,
	}

	request, err := http.NewRequest("GET", API_URL, nil)
	request.Header.Add("Authorization", "Basic " + BasicAuth(CONSUMER_KEY, CONSUMER_SECRET))
	response, err := client.Do(request)

	return json.Marshal(response)
}

//Since Golang drops specified headers on redirects
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.SetBasicAuth(CONSUMER_KEY, CONSUMER_SECRET)
	return nil
}


func BasicAuth(consumerKey, consumerSecret string)  string{
	auth := CONSUMER_KEY + ":" + CONSUMER_SECRET
	return base64.StdEncoding.EncodeToString([] byte(auth))

}