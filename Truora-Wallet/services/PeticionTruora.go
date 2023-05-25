package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

//peticion post a truora

func Peticion(dni, country string) models.Response {
	client := &http.Client{}
	dir := "https://api.checks.truora.com/v1/checks"
	Truora_key := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k2N2RhMjFhODc4MzIxZDBiZGY0ZDdhNDYwOTc3MDgwMCIsImV4cCI6MzI2MTc0NDI2NiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ5NDQyNjYsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xXzhmNVk3OEsyUyIsImp0aSI6ImZkYmE2YTQ3LTYxNDEtNDA1Ni1hMWU5LTM5OTQ3NWRlZDc3ZiIsImtleV9uYW1lIjoicHJ1ZWJhIiwia2V5X3R5cGUiOiJiYWNrZW5kIiwidXNlcm5hbWUiOiJjb3JyZW91bml2YWxsZXZhbGVudGluYWNvYm8tcHJ1ZWJhIn0.ivni00IjVb9e4Um420xHol-M9z4-sugTxE8L0_qIZ8Q`

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

	return response
}
