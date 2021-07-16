package users

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

type User struct {
	Id             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

const usersFilePath = "../source_data/users.json"

func LoadUsers(testFilePath string) ([]User, error) {
	//open the files
	filePath := usersFilePath
	if testFilePath != "" {
		filePath = testFilePath
	}
	usersFilePtr, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() (err error) {
		// ensure we close resource
		if rErr := usersFilePtr.Close(); rErr != nil {
			return rErr
		}
		return nil
	}()

	decoder := json.NewDecoder(usersFilePtr)
	var users []User
	err = decoder.Decode(&users)
	if err != nil && err != io.EOF {
		// return error when the JSON is invalid and not empty
		return nil, err
	}
	return users, nil
}

func SearchUsers(users []User, ident string, value string) (userList []User) {
	for _, user := range users {
		switch ident {
		case "_id":
			if strconv.Itoa(user.Id) == value {
				userList = append(userList, user)
				return
			}
		case "url":
			if user.URL == value {
				userList = append(userList, user)
				return
			}
		case "external_id":
			if user.ExternalId == value {
				userList = append(userList, user)
				return
			}
		case "name":
			if user.Name == value {
				userList = append(userList, user)
				return
			}
		case "alias":
			if user.Alias == value {
				userList = append(userList, user)
				return
			}
		case "created_at":
			if user.CreatedAt == value {
				userList = append(userList, user)
				return
			}
		case "active":
			if strconv.FormatBool(user.Active) == value {
				userList = append(userList, user)
				return
			}
		case "verified":
			if strconv.FormatBool(user.Verified) == value {
				userList = append(userList, user)
				return
			}
		case "shared":
			if strconv.FormatBool(user.Shared) == value {
				userList = append(userList, user)
				return
			}
		case "locale":
			if user.Locale == value {
				userList = append(userList, user)
				return
			}
		case "timezone":
			if user.Timezone == value {
				userList = append(userList, user)
				return
			}
		case "last_login_at":
			if user.LastLoginAt == value {
				userList = append(userList, user)
				return
			}
		case "email":
			if user.Email == value {
				userList = append(userList, user)
				return
			}
		case "phone":
			if user.Phone == value {
				userList = append(userList, user)
				return
			}
		case "signature":
			if user.Signature == value {
				userList = append(userList, user)
				return
			}
		case "organization_id":
			if strconv.Itoa(user.OrganizationId) == value {
				userList = append(userList, user)
			}
		case "tags":
			for _, tag := range user.Tags {
				if tag == value {
					userList = append(userList, user)
					return
				}
			}
		case "suspended":
			if strconv.FormatBool(user.Suspended) == value {
				userList = append(userList, user)
				return
			}
		case "role":
			if user.Role == value {
				userList = append(userList, user)
				return
			}
		default:
			return
		}
	}
	return
}
