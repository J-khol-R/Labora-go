package models

import "time"

type Response struct {
	Check struct {
		CheckID        string `json:"check_id"`
		CompanySummary struct {
			CompanyStatus string `json:"company_status"`
			Result        string `json:"result"`
		} `json:"company_summary"`
		Country      string    `json:"country"`
		CreationDate time.Time `json:"creation_date"`
		NameScore    int       `json:"name_score"`
		IDScore      int       `json:"id_score"`
		Score        int       `json:"score"`
		Scores       []struct {
			DataSet  string `json:"data_set"`
			Severity string `json:"severity"`
			Score    int    `json:"score"`
			Result   string `json:"result"`
			ByID     struct {
				Result   string `json:"result"`
				Score    int    `json:"score"`
				Severity string `json:"severity"`
			} `json:"by_id"`
			ByName struct {
				Result   string `json:"result"`
				Score    int    `json:"score"`
				Severity string `json:"severity"`
			} `json:"by_name"`
		} `json:"scores"`
		Status   string `json:"status"`
		Statuses []struct {
			DatabaseID    string   `json:"database_id"`
			DatabaseName  string   `json:"database_name"`
			Status        string   `json:"status"`
			InvalidInputs []string `json:"invalid_inputs,omitempty"`
			DataSet       string   `json:"data_set,omitempty"`
		} `json:"statuses"`
		Summary struct {
			IdentityStatus string `json:"identity_status"`
			NamesFound     []struct {
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Count     int    `json:"count"`
			} `json:"names_found"`
			Result string `json:"result"`
		} `json:"summary"`
		UpdateDate     time.Time `json:"update_date"`
		VehicleSummary struct {
			Result        string `json:"result"`
			VehicleStatus string `json:"vehicle_status"`
		} `json:"vehicle_summary"`
		NationalID string `json:"national_id"`
		Type       string `json:"type"`
	} `json:"check"`
	Details string `json:"details"`
	Self    string `json:"self"`
}
