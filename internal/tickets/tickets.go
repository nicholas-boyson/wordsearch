package tickets

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

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

const ticketFilePath = "../source_data/tickets.json"

func LoadTickets(testFilePath string) ([]Ticket, error) {
	//open the files
	filePath := ticketFilePath
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

func SearchTickets(tickets []Ticket, ident string, value string) (ticketList []Ticket) {
	for _, ticket := range tickets {
		switch ident {
		case "_id":
			if ticket.Id == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "url":
			if ticket.URL == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "external_id":
			if ticket.ExternalId == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "created_at":
			if ticket.CreatedAt == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "type":
			if ticket.Type == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "subject":
			if ticket.Subject == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "description":
			if ticket.Description == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "priority":
			if ticket.Priority == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "status":
			if ticket.Status == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "submitter_id":
			if strconv.Itoa(ticket.SubmitterId) == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "assignee_id":
			if strconv.Itoa(ticket.AssigneeId) == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "organization_id":
			if strconv.Itoa(ticket.OrganizationId) == value {
				ticketList = append(ticketList, ticket)
			}
		case "tags":
			for _, tag := range ticket.Tags {
				if tag == value {
					ticketList = append(ticketList, ticket)
					return
				}
			}
		case "has_incidents":
			if strconv.FormatBool(ticket.HasIncidents) == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "due_at":
			if ticket.DueAt == value {
				ticketList = append(ticketList, ticket)
				return
			}
		case "via":
			if ticket.Via == value {
				ticketList = append(ticketList, ticket)
				return
			}
		default:
			return
		}
	}
	return
}
