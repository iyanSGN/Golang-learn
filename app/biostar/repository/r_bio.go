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


	//GET USER
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
	req.Header.Set("bs-session-id", "da1b31d26beb4705862383b4ff0ee60c")
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

	//POST USER
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

func PostUser(post Users) (Users, error) {
	url := "https://192.168.88.79:8443/api/users" // Replace with your actual API endpoint
	method := "POST"
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	// Create a JSON payload using a struct and marshaling it
	payload := Users{
		User{
			UserID: post.User.UserID,
			UserGroupID: post.User.UserGroupID,
			StartDatetime: "2001-01-01T00:00:00.00Z",
			ExpiryDatetime: "2030-12-31T23:59:00.00Z",
			Name: post.User.Name,
			Email: post.User.Email,
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
	req.Header.Set("bs-session-id", "ee5117a427bb4f7f9afc135d7d5beb75")

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

	var users Users
	if err := json.Unmarshal(body, &users);
	err != nil {
		fmt.Println("error decoding JSON: ", err.Error())
		return post, nil
	}

	return post, nil
}

	//DELETE USER
func DeleteUser() error {
	url := "https://192.168.88.79:8443/api/users?id=875"
	method := "DELETE"

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

		client := &http.Client{Transport: transport}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

		req.Header.Set("Content-Type", "application/json") // Set the content type
		req.Header.Set("bs-session-id", "ee5117a427bb4f7f9afc135d7d5beb75")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))

	return nil
}


	//LOGIN ADMIN
	type Admin struct {	
		LoginID			string		`json:"login_id"`
		Password		string		`json:"password"`
	}
	
	type Login	struct {
		Admin			Admin		`json:"User"`
	}

	func LoginAdmin(post Login) (Login, error) {
		url := "https://192.168.88.79:8443/api/login" // Replace with your actual API endpoint
		method := "POST"
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	
		client := &http.Client{Transport: transport}
	
		// Create a JSON payload using a struct and marshaling it
		payload := Login{
			Admin{
				LoginID: post.Admin.LoginID,
				Password: post.Admin.Password,
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
		req.Header.Set("bs-session-id", "da1b31d26beb4705862383b4ff0ee60c")
	
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
	
		var users Users
		if err := json.Unmarshal(body, &users);
		err != nil {
			fmt.Println("error decoding JSON: ", err.Error())
			return post, nil
		}
	
		return post, nil
	}



	//Logout Admin
	func LogoutAdmin() (string, error) {
		apiUrl := "https://192.168.88.79:8443/api/logout"
	
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	
		client := &http.Client{Transport: transport}
	
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return "request", nil
		}
	
		// req.Header.Set("Host", "192.168.88.79:8443")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("bs-session-id", "de078920afbc424f9e39aeb155320ee4")
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
			return "response", nil
		}
	
		var users APIResponse
		if err := json.Unmarshal(body, &users); err != nil {
			fmt.Println("error decoding JSON:", err)
			return "decoding", nil
		}

		return "success", nil
	
	}

