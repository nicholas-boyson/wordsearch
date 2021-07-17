package users

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

//User defines the user
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

const usersFilePath = "internal/source_data/users.json"

// LoadUsers process to load the users datastore into a slice
func LoadUsers(testFilePath string) ([]User, error) {
	//open the files
	absPath, err := filepath.Abs(usersFilePath)
	if err != nil {
		return nil, err
	}
	filePath := absPath
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

//SearchUsers return slice of users that match provided ident and value
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

// ValidSearchTerms checks an ident against a list of valid options and returns true if it exists
func ValidSearchTerms(ident string) bool {
	validIdents := []string{"_id", "url", "external_id", "name", "alias", "created_at", "active", "verified", "shared", "locale", "timezone", "last_login_at", "email", "phone", "signature", "organization_id", "tags", "suspended", "role"}
	for _, v := range validIdents {
		if v == ident {
			return true
		}
	}
	return false
}
