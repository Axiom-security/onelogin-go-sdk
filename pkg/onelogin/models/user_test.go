package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestUserJSONMarshaling() {
	baseTime := time.Time{} // zero time
	tests := []struct {
		name     string
		user     User
		expected string
		wantErr  bool
	}{
		{
			name: "full user object",
			user: User{
				ID:                123,
				Firstname:         "John",
				Lastname:          "Doe",
				Email:             "john.doe@example.com",
				Username:          "johndoe",
				State:             StateApproved,
				Status:            StatusActive,
				GroupID:           789,
				DirectoryID:       456,
				ExternalID:        12345,
				CustomAttributes:  map[string]interface{}{"department": "IT"},
				CreatedAt:         baseTime,
				UpdatedAt:         baseTime,
				ActivatedAt:       baseTime,
				LastLogin:         baseTime,
				PasswordChangedAt: baseTime,
				LockedUntil:       baseTime,
				InvitationSentAt:  baseTime,
			},
			expected: `{
				"firstname": "John",
				"lastname": "Doe",
				"username": "johndoe",
				"email": "john.doe@example.com",
				"state": 1,
				"status": 1,
				"group_id": 789,
				"directory_id": 456,
				"external_id": 12345,
				"id": 123,
				"custom_attributes": {"department": "IT"},
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z",
				"activated_at": "0001-01-01T00:00:00Z",
				"last_login": "0001-01-01T00:00:00Z",
				"password_changed_at": "0001-01-01T00:00:00Z",
				"locked_until": "0001-01-01T00:00:00Z",
				"invitation_sent_at": "0001-01-01T00:00:00Z"
			}`,
			wantErr: false,
		},
		{
			name: "minimal user object",
			user: User{
				ID:                456,
				Username:          "janedoe",
				Email:             "jane.doe@example.com",
				CreatedAt:         baseTime,
				UpdatedAt:         baseTime,
				ActivatedAt:       baseTime,
				LastLogin:         baseTime,
				PasswordChangedAt: baseTime,
				LockedUntil:       baseTime,
				InvitationSentAt:  baseTime,
			},
			expected: `{
				"username": "janedoe",
				"email": "jane.doe@example.com",
				"id": 456,
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z",
				"activated_at": "0001-01-01T00:00:00Z",
				"last_login": "0001-01-01T00:00:00Z",
				"password_changed_at": "0001-01-01T00:00:00Z",
				"locked_until": "0001-01-01T00:00:00Z",
				"invitation_sent_at": "0001-01-01T00:00:00Z"
			}`,
			wantErr: false,
		},
		{
			name: "boundary values",
			user: User{
				ID:                2147483647, // max int32
				State:             StateApproved,
				Status:            StatusActive,
				GroupID:           9223372036854775807, // max int64
				DirectoryID:       9223372036854775807,
				ExternalID:        9223372036854775807,
				CreatedAt:         baseTime,
				UpdatedAt:         baseTime,
				ActivatedAt:       baseTime,
				LastLogin:         baseTime,
				PasswordChangedAt: baseTime,
				LockedUntil:       baseTime,
				InvitationSentAt:  baseTime,
			},
			expected: `{
				"id": 2147483647,
				"state": 1,
				"status": 1,
				"group_id": 9223372036854775807,
				"directory_id": 9223372036854775807,
				"external_id": 9223372036854775807,
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z",
				"activated_at": "0001-01-01T00:00:00Z",
				"last_login": "0001-01-01T00:00:00Z",
				"password_changed_at": "0001-01-01T00:00:00Z",
				"locked_until": "0001-01-01T00:00:00Z",
				"invitation_sent_at": "0001-01-01T00:00:00Z"
			}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			actual, err := json.Marshal(tt.user)
			if tt.wantErr {
				assert.Error(s.T(), err)
				return
			}

			assert.NoError(s.T(), err)

			var expectedJSON, actualJSON map[string]interface{}
			err = json.Unmarshal([]byte(tt.expected), &expectedJSON)
			assert.NoError(s.T(), err)

			err = json.Unmarshal(actual, &actualJSON)
			assert.NoError(s.T(), err)

			assert.Equal(s.T(), expectedJSON, actualJSON)
		})
	}
}

func (s *UserTestSuite) TestUserAppJSONMarshaling() {
	tests := []struct {
		name     string
		userApp  UserApp
		expected string
		wantErr  bool
	}{
		{
			name: "full user app object",
			userApp: UserApp{
				ID:                  intPtr(123),
				IconURL:             strPtr("https://example.com/icon.png"),
				LoginID:             int64Ptr(456),
				ProvisioningStatus:  strPtr("active"),
				ProvisioningState:   strPtr("enabled"),
				ProvisioningEnabled: boolPtr(true),
			},
			expected: `{
				"id": 123,
				"icon_url": "https://example.com/icon.png",
				"login_id": 456,
				"provisioning_status": "active",
				"provisioning_state": "enabled",
				"provisioning_enabled": true
			}`,
			wantErr: false,
		},
		{
			name: "boundary values",
			userApp: UserApp{
				ID:      intPtr(2147483647),            // max int32
				LoginID: int64Ptr(9223372036854775807), // max int64
			},
			expected: `{
				"id": 2147483647,
				"login_id": 9223372036854775807
			}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			actual, err := json.Marshal(tt.userApp)
			if tt.wantErr {
				assert.Error(s.T(), err)
				return
			}

			assert.NoError(s.T(), err)

			var expectedJSON, actualJSON map[string]interface{}
			err = json.Unmarshal([]byte(tt.expected), &expectedJSON)
			assert.NoError(s.T(), err)

			err = json.Unmarshal(actual, &actualJSON)
			assert.NoError(s.T(), err)

			assert.Equal(s.T(), expectedJSON, actualJSON)
		})
	}
}

func (s *UserTestSuite) TestUserQueryValidation() {
	now := time.Now()
	tests := []struct {
		name      string
		query     UserQuery
		validKeys []string
		expectErr bool
	}{
		{
			name: "valid query with strings",
			query: UserQuery{
				Email:    strPtr("test@example.com"),
				Username: strPtr("testuser"),
			},
			validKeys: []string{"email", "username"},
			expectErr: false,
		},
		{
			name: "valid query with dates",
			query: UserQuery{
				CreatedSince: timePtr(now.Add(-24 * time.Hour)),
				CreatedUntil: timePtr(now),
			},
			validKeys: []string{"createdSince", "createdUntil"},
			expectErr: false,
		},
		{
			name: "invalid empty string",
			query: UserQuery{
				Email: strPtr(""),
			},

			validKeys: []string{"email"},
			expectErr: true,
		},
		{
			name: "invalid zero time",
			query: UserQuery{
				CreatedSince: timePtr(time.Time{}),
			},
			validKeys: []string{"createdSince"},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			validators := tt.query.GetKeyValidators()

			for _, key := range tt.validKeys {
				validator, exists := validators[key]
				assert.True(s.T(), exists, "Validator not found for key %s", key)

				value := getFieldValue(&tt.query, key)
				assert.NotNil(s.T(), value, "Value not found for key %s", key)

				isValid := validator(value)
				if tt.expectErr {
					assert.False(s.T(), isValid, "Expected validation to fail for %s", key)
				} else {
					assert.True(s.T(), isValid, "Expected validation to pass for %s", key)
				}
			}
		})
	}
}

// Helper functions for creating pointers
func strPtr(s string) *string {
	return &s
}

func intPtr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func timePtr(t time.Time) *time.Time {
	return &t
}

// Helper function to get field value by name
func getFieldValue(query *UserQuery, fieldName string) interface{} {
	switch fieldName {
	case "email":
		return query.Email
	case "username":
		return query.Username
	case "firstname":
		return query.Firstname
	case "lastname":
		return query.Lastname
	case "samaccountname":
		return query.Samaccountname
	case "directoryID":
		return query.DirectoryID
	case "externalID":
		return query.ExternalID
	case "appID":
		return query.AppID
	case "userIDs":
		return query.UserIDs
	case "fields":
		return query.Fields
	case "createdSince":
		return query.CreatedSince
	case "createdUntil":
		return query.CreatedUntil
	case "updatedSince":
		return query.UpdatedSince
	case "updatedUntil":
		return query.UpdatedUntil
	case "lastLoginSince":
		return query.LastLoginSince
	case "lastLoginUntil":
		return query.LastLoginUntil
	case "limit":
		return query.Limit
	case "page":
		return query.Page
	case "cursor":
		return query.Cursor
	default:
		return nil
	}
}
