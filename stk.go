package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	_ "net/http/cookiejar"
)


const CONSUMER_KEY = "consumerkey"
const CONSUMER_SECRET = "consumser_secret"
const API_URL = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"


//A struct response to map the response to JSON
type Response struct {

}

func getMPESAToken() (interface{}, error) {
	client := &http.Client{
		Jar: http.CookieJar(nil),
		CheckRedirect: redirectPolicyFunc,
	}

	request, _ := http.NewRequest("GET", API_URL, nil)
	request.Header.Add("Authorization", "Basic " + BasicAuth(CONSUMER_KEY, CONSUMER_SECRET))
	response, _ := client.Do(request)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	return json.Marshal(responseData)
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

//Function to perform the STK Push
func STKPush()  {
	response, err := getMPESAToken()
	if err != nil{
	 panic(err)
	}
	accessToken := response['access_token']

	var encoded_data []string
	passkey := base64.Encoding{}
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

	//Adding access token to request headers
	headers := map[string] string{"Authorization":"Bearer " + accessToken}

}


func main(){


}