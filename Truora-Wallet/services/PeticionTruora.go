package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	"github.com/joho/godotenv"
)

var client = &http.Client{}

const (
	dir = "https://api.checks.truora.com/v1/checks"
)

func SendRequestToTruora(dni, country string) (models.Response, error) {
	err := godotenv.Load()
	if err != nil {
		return models.Response{}, err
	}
	Truora_key := os.Getenv("TRUORA_KEY")
	formData := url.Values{
		"national_id":     {dni},
		"country":         {country},
		"type":            {"person"},
		"user_authorized": {"true"},
	}

	req, err := http.NewRequest("POST", dir, strings.NewReader(formData.Encode()))
	if err != nil {
		return models.Response{}, err
	}

	req.Header.Add("Truora-API-Key", Truora_key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return models.Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Response{}, err
	}

	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.Response{}, err
	}

	if response.Check.Score == -1 {
		time.Sleep(4 * time.Second)
		peticion, err := PeticionGet(response.Check.CheckID)
		if err != nil {
			return models.Response{}, err
		}
		return peticion, nil
	}
	return response, nil
}

func PeticionGet(check_id string) (models.Response, error) {
	err := godotenv.Load()
	if err != nil {
		return models.Response{}, err
	}
	Truora_key := os.Getenv("TRUORA_KEY")
	req, err := http.NewRequest("GET", dir+"/"+check_id, nil)
	if err != nil {
		return models.Response{}, err
	}

	req.Header.Add("Truora-API-Key", Truora_key)

	resp, err := client.Do(req)
	if err != nil {
		return models.Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Response{}, err
	}

	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.Response{}, err
	}
	return response, nil
}
