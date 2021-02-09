package main

import (
	"encoding/base64"
	"net/http"
)


const CONSUMER_KEY = "consumerkey"
const CONSUMER_SECRET = "consumser_secret"
const API_URL = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

func main(){

}

func getMPESAToken()  {


	request, err := http.NewRequest("GET", API_URL, nil)
	request.Header.Add("Authorization", "Basic " + BasicAuth(CONSUMER_KEY, CONSUMER_SECRET))
	//response, err := cl


}

func redirectPolicyFUnc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic " + BasicAuth(CONSUMER_KEY, CONSUMER_SECRET))
}

func BasicAuth(consumerKey, consumerSecret string)  string{
	auth := CONSUMER_KEY + ":" + CONSUMER_SECRET
	return base64.StdEncoding.EncodeToString([] byte(auth))

}