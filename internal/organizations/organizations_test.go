package organizations

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationsPath(t *testing.T) {
	assert.Equal(t, "../source_data/organizations.json", organizationsFilePath)
}

func TestLoadOrganizations(t *testing.T) {
	tests := []struct {
		test          string
		testFilePath  string
		organizations []Organization
		err           error
	}{
		{
			test:         "MissingOrganizationsFile",
			testFilePath: "missingOrganizationsFile",
			err:          errors.New("open missingOrganizationsFile: The system cannot find the file specified."),
		},
		{
			test:         "InvalidOrganizationsFile1",
			testFilePath: "test_files/invalid_json_organizations.json",
			err:          errors.New("json: cannot unmarshal string into Go struct field Organization._id of type int"),
		},
		{
			test:         "GoodOrganizationsFile",
			testFilePath: "test_files/good_organizations.json",
			organizations: []Organization{
				{
					Id:            101,
					URL:           "http://initech.zendesk.com/api/v2/organizations/101.json",
					ExternalId:    "9270ed79-35eb-4a38-a46f-35725197ea8d",
					Name:          "Enthaze",
					DomainNames:   []string{"kage.com", "ecratic.com", "endipin.com", "zentix.com"},
					CreatedAt:     "2016-05-21T11:10:28 -10:00",
					Details:       "MegaCorp",
					SharedTickets: false,
					Tags:          []string{"Fulton", "West", "Rodriguez", "Farley"},
				},
				{
					Id:            102,
					URL:           "http://initech.zendesk.com/api/v2/organizations/102.json",
					ExternalId:    "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
					Name:          "Nutralab",
					DomainNames:   []string{"trollery.com", "datagen.com", "bluegrain.com", "dadabase.com"},
					CreatedAt:     "2016-04-07T08:21:44 -10:00",
					Details:       "Non profit",
					SharedTickets: false,
					Tags:          []string{"Cherry", "Collier", "Fuentes", "Trevino"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			organizations, err := LoadOrganizations(tt.testFilePath)
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.organizations, organizations)
		})
	}
}

func BenchmarkOrganisationLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := LoadOrganizations("")
		assert.Nil(b, err)
	}
}

func BenchmarkOrganisationLoadMapField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := LoadOrganizationsMapField("_id")
		assert.Nil(b, err)
	}
}

func BenchmarkOrganisationLoadThenMap(b *testing.B) {
	orgs, err := LoadOrganizations("")
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		_ = OrganizationsMapField(orgs, "_id")
	}
}

func BenchmarkOrganisationLoadThenSearch(b *testing.B) {
	orgs, err := LoadOrganizations("")
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		_ = SearchOrganizations(orgs, "_id", "125")
	}
}
