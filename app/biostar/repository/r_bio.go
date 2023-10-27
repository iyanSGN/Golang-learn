package repository

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "time"
)

type Rows struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	LoginID string `json:"login_id"`
}

type UserCollection struct {
	Total string `json:"total"`
	Rows  []Rows `json:"rows"`
}

type APIResponse struct {
	UserCollection UserCollection `json:"UserCollection"`
}

type getApiResponse struct {
	UserCollection UserCollection `json:"UserCollection"`
}

func GetUser() ([]getApiResponse, error) {
	apiUrl := "https://192.168.88.79:8443/api/users"

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// req.Header.Set("Host", "192.168.88.79:8443")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("bs-session-id", "440ac596d7f743d2aac948e031bf3e7e")
	// req.Header.Set("Cookie", "bs-session-id=8a2efee39aa44023b89e03b48ab107e0")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return nil, err
	}

	var users APIResponse
	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Println("error decoding JSON:", err)
		return nil, err
	}

	getusers := []getApiResponse{{UserCollection: users.UserCollection}}

	return getusers, nil

}


type User struct {	
	UserID     		string 			`json:"user_id"`
	UserGroupID    struct {
		ID string `json:"id"`
	} `json:"user_group_id"`
	StartDatetime  	string 		`json:"start_datetime"`
	ExpiryDatetime  string 		`json:"expiry_datetime"`
	Name            string 			`json:"name"`
	Email           string 			`json:"email"`
}

type Users	struct {
	User			User		`json:"User"`
}

type PostUsers struct {
	UserID     		string 			`json:"user_id"`
	UserGroupID    struct {
		ID string `json:"id"`
	} `json:"user_group_id"`
	StartDatetime   string 		`json:"start_datetime"`
	ExpiryDatetime  string 		`json:"expiry_datetime"`
	Name            string 			`json:"name"`
	Email           string 			`json:"email"`
}

type Post struct {
	PostUsers		PostUsers		`json:"post"`
}

func PostUser(post Post) (Post, error) {
	url := "https://192.168.88.79:8443/api/users" // Replace with your actual API endpoint
	method := "POST"
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	// Create a JSON payload using a struct and marshaling it
	payload := Users{
		User{
			UserID: "873",
			UserGroupID: struct {
				ID string `json:"id"`
			}{
				ID: "1",
			},
			StartDatetime: "2001-01-01T00:00:00.00Z",
			ExpiryDatetime: "2030-12-31T23:59:00.00Z",
			Name: "iyan",
			Email: "iyan040@gmail.com",
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(payloadBytes))
	
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type
	req.Header.Set("bs-session-id", "a7fc8ced292f45bc85102103cb3fcf1a")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Println("error req: %w", err))
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(fmt.Println("error body: %w", err))
	}

	fmt.Println("Response Body:", string(body))

	var users User
	if err := json.Unmarshal(body, &users);
	err != nil {
		fmt.Println("error decoding JSON: ", err.Error())
		return post, nil
	}

	return post, nil
}
