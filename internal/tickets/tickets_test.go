package tickets

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketsPath(t *testing.T) {
	assert.Equal(t, "internal/source_data/tickets.json", ticketFilePath)
}

func TestLoadTickets(t *testing.T) {
	tests := []struct {
		test         string
		testFilePath string
		tickets      []Ticket
		err          error
	}{
		{
			test:         "MissingTicketFile",
			testFilePath: "missingTicketFile",
			err:          errors.New("open missingTicketFile: The system cannot find the file specified."),
		},
		{
			test:         "InvalidTicketFile1",
			testFilePath: "test_files/invalid_json_tickets.json",
			err:          errors.New("json: cannot unmarshal string into Go struct field Ticket.has_incidents of type bool"),
		},
		{
			test:         "GoodTicketFile",
			testFilePath: "test_files/good_tickets.json",
			tickets: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			tickets, err := LoadTickets(tt.testFilePath)
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.tickets, tickets)
		})
	}
}

func TestValidSearchTerms(t *testing.T) {
	tests := []struct {
		test       string
		searchTerm string
		result     bool
	}{
		{
			test:       "ValidTerm",
			searchTerm: "_id",
			result:     true,
		},
		{
			test:       "InvalidTerm",
			searchTerm: "InvalidTerm",
			result:     false,
		},
	}

	for _, tt := range tests {
		result := ValidSearchTerms(tt.searchTerm)
		assert.Equal(t, tt.result, result)
	}
}

func TestSearchTickets(t *testing.T) {
	tests := []struct {
		test   string
		ident  string
		value  string
		input  []Ticket
		result []Ticket
	}{
		{
			test:  "IntIdent",
			ident: "submitter_id",
			value: "38",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
			},
		},
		{
			test:  "BoolIdent",
			ident: "has_incidents",
			value: "false",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
		},
		{
			test:  "ArrayIdent",
			ident: "tags",
			value: "Ohio",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
			},
		},
		{
			test:  "StringIdent",
			ident: "url",
			value: "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: []Ticket{
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
		},
		{
			test:  "MultipleResults",
			ident: "type",
			value: "incident",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
		},
		{
			test:  "NoMatch",
			ident: "_id",
			value: "103",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Micronesia",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: nil,
		},
		{
			test:  "EmptyValueMatch",
			ident: "subject",
			value: "",
			input: []Ticket{
				{
					Id:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					ExternalId:     "9210cdc9-4bee-485f-a078-35396cd74063",
					CreatedAt:      "2016-04-28T11:19:34 -10:00",
					Type:           "incident",
					Subject:        "A Catastrophe in Korea (North)",
					Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					Priority:       "high",
					Status:         "pending",
					SubmitterId:    38,
					AssigneeId:     24,
					OrganizationId: 116,
					Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
					HasIncidents:   false,
					DueAt:          "2016-07-31T02:37:50 -10:00",
					Via:            "web",
				},
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
			result: []Ticket{
				{
					Id:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
					URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
					ExternalId:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
					CreatedAt:      "2016-04-14T08:32:31 -10:00",
					Type:           "incident",
					Subject:        "",
					Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
					Priority:       "low",
					Status:         "hold",
					SubmitterId:    71,
					AssigneeId:     38,
					OrganizationId: 112,
					Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
					HasIncidents:   false,
					DueAt:          "2016-08-15T05:37:32 -10:00",
					Via:            "chat",
				},
			},
		},
	}

	for _, tt := range tests {
		result := SearchTickets(tt.input, tt.ident, tt.value)
		assert.Equal(t, tt.result, result)
	}
}

func BenchmarkTicketsLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := LoadTickets("")
		assert.Nil(b, err)
	}
}

func BenchmarkTicketsLoadThenSearch(b *testing.B) {
	tickets, err := LoadTickets("")
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		_ = SearchTickets(tickets, "organization_id", "125")
	}
}
