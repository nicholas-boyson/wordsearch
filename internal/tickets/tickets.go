package tickets

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// Ticket defines the ticket
type Ticket struct {
	Id             string   `json:"_id"`
	URL            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterId    int      `json:"submitter_id"`
	AssigneeId     int      `json:"assignee_id"`
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

const ticketFilePath = "internal/source_data/tickets.json"

// LoadTickets process to load the tickets datastore into a slice
func LoadTickets(testFilePath string) ([]Ticket, error) {
	//open the files
	absPath, err := filepath.Abs(ticketFilePath)
	if err != nil {
		return nil, err
	}
	filePath := absPath
	if testFilePath != "" {
		filePath = testFilePath
	}
	ticketsFilePtr, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() (err error) {
		// ensure we close resource
		if rErr := ticketsFilePtr.Close(); rErr != nil {
			return rErr
		}
		return nil
	}()

	decoder := json.NewDecoder(ticketsFilePtr)
	var tickets []Ticket
	err = decoder.Decode(&tickets)
	if err != nil && err != io.EOF {
		// return error when the JSON is invalid and not empty
		return nil, err
	}
	return tickets, nil
}

//SearchTickets return slice of tickets that match provided ident and value
func SearchTickets(tickets []Ticket, ident string, value string) (ticketList []Ticket) {
	for _, ticket := range tickets {
		switch ident {
		case "_id":
			if ticket.Id == value {
				ticketList = append(ticketList, ticket)
			}
		case "url":
			if ticket.URL == value {
				ticketList = append(ticketList, ticket)
			}
		case "external_id":
			if ticket.ExternalId == value {
				ticketList = append(ticketList, ticket)
			}
		case "created_at":
			if ticket.CreatedAt == value {
				ticketList = append(ticketList, ticket)
			}
		case "type":
			if ticket.Type == value {
				ticketList = append(ticketList, ticket)
			}
		case "subject":
			if ticket.Subject == value {
				ticketList = append(ticketList, ticket)
			}
		case "description":
			if ticket.Description == value {
				ticketList = append(ticketList, ticket)
			}
		case "priority":
			if ticket.Priority == value {
				ticketList = append(ticketList, ticket)
			}
		case "status":
			if ticket.Status == value {
				ticketList = append(ticketList, ticket)
			}
		case "submitter_id":
			if strconv.Itoa(ticket.SubmitterId) == value {
				ticketList = append(ticketList, ticket)
			}
		case "assignee_id":
			if strconv.Itoa(ticket.AssigneeId) == value {
				ticketList = append(ticketList, ticket)
			}
		case "organization_id":
			if strconv.Itoa(ticket.OrganizationId) == value {
				ticketList = append(ticketList, ticket)
			}
		case "tags":
			for _, tag := range ticket.Tags {
				if tag == value {
					ticketList = append(ticketList, ticket)
				}
			}
		case "has_incidents":
			if strconv.FormatBool(ticket.HasIncidents) == value {
				ticketList = append(ticketList, ticket)
			}
		case "due_at":
			if ticket.DueAt == value {
				ticketList = append(ticketList, ticket)
			}
		case "via":
			if ticket.Via == value {
				ticketList = append(ticketList, ticket)
			}
		default:
			return
		}
	}
	return
}

// ValidSearchTerms checks an ident against a list of valid options and returns true if it exists
func ValidSearchTerms(ident string) bool {
	validIdents := []string{"_id", "url", "external_id", "created_at", "type", "subject", "description", "priority", "status", "submitter_id", "assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via"}
	for _, v := range validIdents {
		if v == ident {
			return true
		}
	}
	return false
}
