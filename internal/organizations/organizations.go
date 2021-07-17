package organizations

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

//Organization struct defining an organization
type Organization struct {
	Id            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalId    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

const organizationsFilePath = "internal/source_data/organizations.json"

// LoadOrganizations process to load the organizations datastore into a slice
func LoadOrganizations(testFilePath string) ([]Organization, error) {
	//open the files
	absPath, err := filepath.Abs(organizationsFilePath)
	if err != nil {
		return nil, err
	}
	filePath := absPath
	if testFilePath != "" {
		filePath = testFilePath
	}
	organizationsFilePtr, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() (err error) {
		// ensure we close resource
		if rErr := organizationsFilePtr.Close(); rErr != nil {
			return rErr
		}
		return nil
	}()

	decoder := json.NewDecoder(organizationsFilePtr)
	var organizations []Organization
	err = decoder.Decode(&organizations)
	if err != nil && err != io.EOF {
		// return error when the JSON is invalid and not empty
		return nil, err
	}
	return organizations, nil
}

//SearchOrganizations function to search over all organizations based on ident and value provided
func SearchOrganizations(organizations []Organization, ident string, value string) Organization {
	for _, org := range organizations {
		switch ident {
		case "_id":
			if strconv.Itoa(org.Id) == value {
				return org
			}
		case "url":
			if org.URL == value {
				return org
			}
		case "external_id":
			if org.ExternalId == value {
				return org
			}
		case "name":
			if org.Name == value {
				return org
			}
		case "domain_names":
			for _, dn := range org.DomainNames {
				if dn == value {
					return org
				}
			}
		case "created_at":
			if org.CreatedAt == value {
				return org
			}
		case "details":
			if org.Details == value {
				return org
			}
		case "shared_tickets":
			if strconv.FormatBool(org.SharedTickets) == value {
				return org
			}
		case "tags":
			for _, tag := range org.Tags {
				if tag == value {
					return org
				}
			}
		default:
			return Organization{}
		}
	}
	return Organization{}
}

// ValidSearchTerms checks an ident against a list of valid options and returns true if it exists
func ValidSearchTerms(ident string) bool {
	validIdents := []string{"_id", "url", "external_id", "name", "domain_names", "created_at", "details", "shared_tickets", "tags"}
	for _, v := range validIdents {
		if v == ident {
			return true
		}
	}
	return false
}
