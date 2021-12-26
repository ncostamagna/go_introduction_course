package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	User User `json:"data"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserCreated struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Job       string     `json:"job,omitempty"`
	Age       int        `json:"age,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

const (
	base = "https://reqres.in/api"
)

func main() {
	req, err := GetReqExample(fmt.Sprintf("%s/users/2", base))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(req))
	fmt.Println()
	req, err = Get(fmt.Sprintf("%s/users/2", base))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(req))
	fmt.Println()
	user, err := GetUser("2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	fmt.Println("ID: ", user.ID)
	fmt.Println("Email: ", user.Email)
	fmt.Println("FirstName: ", user.FirstName)
	fmt.Println("LastName: ", user.LastName)

	fmt.Println()
	u, err := CreateUser("nahuel", "engineer")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	fmt.Println("ID: ", u.ID)
	fmt.Println("Name: ", u.Name)
	fmt.Println("Job: ", u.Job)
	fmt.Println("Date: ", u.CreatedAt)

	v, _ := json.Marshal(u)
	fmt.Println(string(v))
}

func GetReqExample(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		fmt.Println("entra")
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func GetUser(userID string) (*User, error) {
	req, err := Get(fmt.Sprintf("%s/users/%s", base, userID))
	if err != nil {
		return nil, err
	}

	var dataResponse Data

	if err := json.Unmarshal(req, &dataResponse); err != nil {
		return nil, err
	}

	return &dataResponse.User, nil
}

func Post(url string, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/text")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func CreateUser(name, job string) (*UserCreated, error) {

	user := &UserCreated{
		Name: name,
		Job:  job,
	}
	req, err := Post(fmt.Sprintf("%s/users", base), user)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(req, user); err != nil {
		return nil, err
	}
	return user, nil
}
