package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersPath(t *testing.T) {
	assert.Equal(t, "../source_data/users.json", usersFilePath)
}

func TestLoadUsers(t *testing.T) {
	tests := []struct {
		test         string
		testFilePath string
		users        []User
		err          error
	}{
		{
			test:         "MissingUsersFile",
			testFilePath: "missingUsersFile",
			err:          errors.New("open missingUsersFile: The system cannot find the file specified."),
		},
		{
			test:         "InvalidUsersFile1",
			testFilePath: "test_files/invalid_json_Users.json",
			err:          errors.New("json: cannot unmarshal string into Go struct field User.shared of type int"),
		},
		{
			test:         "GoodUsersFile",
			testFilePath: "test_files/good_Users.json",
			users: []User{
				{
					Id:             1,
					URL:            "http://initech.zendesk.com/api/v2/users/1.json",
					ExternalId:     "74341f74-9c79-49d5-9611-87ef9b6eb75f",
					Name:           "Francisca Rasmussen",
					Alias:          "Miss Coffey",
					CreatedAt:      "2016-04-15T05:19:46 -10:00",
					Active:         true,
					Verified:       true,
					Shared:         false,
					Locale:         "en-AU",
					Timezone:       "Sri Lanka",
					LastLoginAt:    "2013-08-04T01:03:27 -10:00",
					Email:          "coffeyrasmussen@flotonic.com",
					Phone:          "8335-422-718",
					Signature:      "Don't Worry Be Happy!",
					OrganizationId: 119,
					Tags:           []string{"Springville", "Sutton", "Hartsville/Hartley", "Diaperville"},
					Suspended:      true,
					Role:           "admin",
				},
				{
					Id:             2,
					URL:            "http://initech.zendesk.com/api/v2/users/2.json",
					ExternalId:     "c9995ea4-ff72-46e0-ab77-dfe0ae1ef6c2",
					Name:           "Cross Barlow",
					Alias:          "Miss Joni",
					CreatedAt:      "2016-06-23T10:31:39 -10:00",
					Active:         true,
					Verified:       true,
					Shared:         false,
					Locale:         "zh-CN",
					Timezone:       "Armenia",
					LastLoginAt:    "2012-04-12T04:03:28 -10:00",
					Email:          "jonibarlow@flotonic.com",
					Phone:          "9575-552-585",
					Signature:      "Don't Worry Be Happy!",
					OrganizationId: 106,
					Tags:           []string{"Foxworth", "Woodlands", "Herlong", "Henrietta"},
					Suspended:      false,
					Role:           "admin",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			users, err := LoadUsers(tt.testFilePath)
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.users, users)
		})
	}
}

func BenchmarkUsersLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := LoadUsers("")
		assert.Nil(b, err)
	}
}

func BenchmarkUsersLoadThenSearch(b *testing.B) {
	users, err := LoadUsers("")
	assert.Nil(b, err)
	for i := 0; i < b.N; i++ {
		_ = SearchUsers(users, "organization_id", "125")
	}
}
