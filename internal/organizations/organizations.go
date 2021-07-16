package organizations

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

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

func LoadOrganizationsMapField(ident string) (map[string]Organization, error) {
	//open the files
	var organizationsFilePtr *os.File
	organizationsFilePtr, err := os.Open("../source_data/organizations.json")
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

	var returnList = map[string]Organization{}
	for _, org := range organizations {
		switch ident {
		case "_id":
			returnList[strconv.Itoa(org.Id)] = org
		case "url":
			returnList[org.URL] = org
		case "external_id":
			returnList[org.ExternalId] = org
		case "name":
			returnList[org.Name] = org
		case "domain_names":
			for _, dn := range org.DomainNames {
				returnList[dn] = org
			}
		case "created_at":
			returnList[org.CreatedAt] = org
		case "details":
			returnList[org.Details] = org
		case "shared_tickets":
			returnList[strconv.FormatBool(org.SharedTickets)] = org
		case "tags":
			for _, tag := range org.Tags {
				returnList[tag] = org
			}
		default:
			returnList[strconv.Itoa(org.Id)] = org
		}
	}
	return returnList, nil
}

const organizationsFilePath = "../source_data/organizations.json"

func LoadOrganizations(testFilePath string) ([]Organization, error) {
	//open the files
	filePath := organizationsFilePath
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

func OrganizationsMapField(organizations []Organization, ident string) map[string]Organization {
	var returnList = map[string]Organization{}
	for _, org := range organizations {
		switch ident {
		case "_id":
			returnList[strconv.Itoa(org.Id)] = org
		case "url":
			returnList[org.URL] = org
		case "external_id":
			returnList[org.ExternalId] = org
		case "name":
			returnList[org.Name] = org
		case "domain_names":
			for _, dn := range org.DomainNames {
				returnList[dn] = org
			}
		case "created_at":
			returnList[org.CreatedAt] = org
		case "details":
			returnList[org.Details] = org
		case "shared_tickets":
			returnList[strconv.FormatBool(org.SharedTickets)] = org
		case "tags":
			for _, tag := range org.Tags {
				returnList[tag] = org
			}
		default:
			returnList[strconv.Itoa(org.Id)] = org
		}
	}
	return returnList
}

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
