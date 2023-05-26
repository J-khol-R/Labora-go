package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

func PeticionPost(dni, country string) models.Response {
	client := &http.Client{}
	dir := "https://api.checks.truora.com/v1/checks"
	Truora_key := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k4YWJkOWE1ZGFmNzM1NGQ1YjVlZjVjYTI4MjJhMjA3OSIsImV4cCI6MzI2MTY4OTIwMiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ4ODkyMDIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX3hUSGxqU1d2RCIsImp0aSI6IjM2YTZiNGJlLTM3NTUtNGQzMC04ZTM0LTNmZDMyOGI3ZDk3NCIsImtleV9uYW1lIjoidHJ1Y29kZSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoidHJ1b3JhdGVhbW5ld3Byb2QtdHJ1Y29kZSJ9.PuE6cS6938PbQz_4qMLySs9dr3fywFqqGdfcF6Suw0U`

	formData := url.Values{
		"national_id":     {dni},
		"country":         {country},
		"type":            {"person"},
		"user_authorized": {"true"},
	}

	req, err := http.NewRequest("POST", dir, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error request")
	}

	req.Header.Add("Truora-API-Key", Truora_key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	if response.Check.Score == -1 {
		time.Sleep(4 * time.Second)
		return PeticionGet(response.Check.CheckID)
	} else {
		return response
	}
}

func PeticionGet(check_id string) models.Response {
	client := &http.Client{}
	dir := "https://api.checks.truora.com/v1/checks/"
	Truora_key := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k4YWJkOWE1ZGFmNzM1NGQ1YjVlZjVjYTI4MjJhMjA3OSIsImV4cCI6MzI2MTY4OTIwMiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ4ODkyMDIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX3hUSGxqU1d2RCIsImp0aSI6IjM2YTZiNGJlLTM3NTUtNGQzMC04ZTM0LTNmZDMyOGI3ZDk3NCIsImtleV9uYW1lIjoidHJ1Y29kZSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoidHJ1b3JhdGVhbW5ld3Byb2QtdHJ1Y29kZSJ9.PuE6cS6938PbQz_4qMLySs9dr3fywFqqGdfcF6Suw0U`

	req, err := http.NewRequest("GET", dir+check_id, nil)
	if err != nil {
		fmt.Println("Error request")
	}

	req.Header.Add("Truora-API-Key", Truora_key)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
