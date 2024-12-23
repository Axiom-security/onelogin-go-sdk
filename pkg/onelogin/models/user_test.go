package models

import (
	"encoding/json"
	"fmt"
	"strconv"
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
				GroupID:           "789",
				DirectoryID:       "456",
				TrustedIDPID:      "101112",
				ManagerADID:       "131415",
				ManagerUserID:     "161718",
				ExternalID:        "12345",
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
				"group_id": "789",
				"directory_id": "456",
				"trusted_idp_id": "101112",
				"manager_ad_id": "131415",
				"manager_user_id": "161718",
				"external_id": "12345",
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
			name: "string IDs with various formats",
			user: User{
				ID:                789,
				GroupID:           "group-123",
				DirectoryID:       "dir-456",
				TrustedIDPID:      "idp-789",
				ManagerADID:       "ad-101112",
				ManagerUserID:     "user-131415",
				ExternalID:        "ext-161718",
				CreatedAt:         baseTime,
				UpdatedAt:         baseTime,
				ActivatedAt:       baseTime,
				LastLogin:         baseTime,
				PasswordChangedAt: baseTime,
				LockedUntil:       baseTime,
				InvitationSentAt:  baseTime,
			},
			expected: `{
				"id": 789,
				"group_id": "group-123",
				"directory_id": "dir-456",
				"trusted_idp_id": "idp-789",
				"manager_ad_id": "ad-101112",
				"manager_user_id": "user-131415",
				"external_id": "ext-161718",
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
			name: "large numeric string IDs",
			user: User{
				ID:                2147483647,            // max int32
				GroupID:           "9223372036854775807", // max int64 as string
				DirectoryID:       "9223372036854775806",
				TrustedIDPID:      "9223372036854775805",
				ManagerADID:       "9223372036854775804",
				ManagerUserID:     "9223372036854775803",
				ExternalID:        "9223372036854775802",
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
				"group_id": "9223372036854775807",
				"directory_id": "9223372036854775806",
				"trusted_idp_id": "9223372036854775805",
				"manager_ad_id": "9223372036854775804",
				"manager_user_id": "9223372036854775803",
				"external_id": "9223372036854775802",
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
				Email:       strPtr("test@example.com"),
				Username:    strPtr("testuser"),
				DirectoryID: strPtr("12345"),
				ExternalID:  strPtr("67890"),
				AppID:       strPtr("54321"),
				UserIDs:     strPtr("1,2,3,4,5"),
			},
			validKeys: []string{"email", "username", "directoryID", "externalID", "appID", "userIDs"},
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
		{
			name: "valid numeric strings",
			query: UserQuery{
				DirectoryID: strPtr("12345"),
				ExternalID:  strPtr("67890"),
				AppID:       strPtr("54321"),
			},
			validKeys: []string{"directoryID", "externalID", "appID"},
			expectErr: false,
		},
		{
			name: "valid comma-separated IDs",
			query: UserQuery{
				UserIDs: strPtr("1,2,3,4,5"),
			},
			validKeys: []string{"userIDs"},
			expectErr: false,
		},
		{
			name: "empty numeric strings",
			query: UserQuery{
				DirectoryID: strPtr(""),
				ExternalID:  strPtr(""),
				AppID:       strPtr(""),
			},
			validKeys: []string{"directoryID", "externalID", "appID"},
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

func (s *UserTestSuite) TestUserQueryStringConversions() {
	tests := []struct {
		name          string
		query         UserQuery
		expected      string
		fieldsToCheck []string
		wantErr       bool
	}{
		{
			name: "valid numeric strings",
			query: UserQuery{
				DirectoryID: strPtr("12345"),
				ExternalID:  strPtr("67890"),
				AppID:       strPtr("54321"),
				UserIDs:     strPtr("1,2,3,4,5"),
			},
			expected: `{
				"directory_id": "12345",
				"external_id": "67890",
				"app_id": "54321",
				"user_ids": "1,2,3,4,5"
			}`,
			fieldsToCheck: []string{"directoryID", "externalID", "appID", "userIDs"},
			wantErr:       false,
		},
		{
			name: "large numeric strings",
			query: UserQuery{
				DirectoryID: strPtr("9223372036854775807"), // max int64
				ExternalID:  strPtr("9223372036854775806"),
				AppID:       strPtr("9223372036854775805"),
			},
			expected: `{
				"directory_id": "9223372036854775807",
				"external_id": "9223372036854775806",
				"app_id": "9223372036854775805"
			}`,
			fieldsToCheck: []string{"directoryID", "externalID", "appID"},
			wantErr:       false,
		},
		{
			name: "non-numeric strings",
			query: UserQuery{
				DirectoryID: strPtr("dir-123"),
				ExternalID:  strPtr("ext-456"),
				AppID:       strPtr("app-789"),
			},
			expected: `{
				"directory_id": "dir-123",
				"external_id": "ext-456",
				"app_id": "app-789"
			}`,
			fieldsToCheck: []string{"directoryID", "externalID", "appID"},
			wantErr:       true, // Changed to true since these are not valid numeric strings
		},
		{
			name: "mixed format IDs list",
			query: UserQuery{
				UserIDs: strPtr("1,abc,3,def-456,5"),
			},
			expected: `{
				"user_ids": "1,abc,3,def-456,5"
			}`,
			fieldsToCheck: []string{"userIDs"},
			wantErr:       false, // UserIDs can contain non-numeric values
		},
		{
			name: "empty strings",
			query: UserQuery{
				DirectoryID: strPtr(""),
				ExternalID:  strPtr(""),
				AppID:       strPtr(""),
				UserIDs:     strPtr(""),
			},
			expected: `{
				"directory_id": "",
				"external_id": "",
				"app_id": "",
				"user_ids": ""
			}`,
			fieldsToCheck: []string{"directoryID", "externalID", "appID", "userIDs"},
			wantErr:       true,
		},
		{
			name: "whitespace strings",
			query: UserQuery{
				DirectoryID: strPtr("   "),
				ExternalID:  strPtr("\t"),
				AppID:       strPtr("\n"),
				UserIDs:     strPtr(" , , "),
			},
			expected: `{
				"directory_id": "   ",
				"external_id": "\t",
				"app_id": "\n",
				"user_ids": " , , "
			}`,
			fieldsToCheck: []string{"directoryID", "externalID", "appID", "userIDs"},
			wantErr:       true,
		},
		{
			name: "invalid numeric strings",
			query: UserQuery{
				DirectoryID: strPtr("9223372036854775808"),  // greater than max int64
				ExternalID:  strPtr("-9223372036854775809"), // less than min int64
				AppID:       strPtr("123.456"),              // decimal
			},
			expected: `{
				"directory_id": "9223372036854775808",
				"external_id": "-9223372036854775809",
				"app_id": "123.456"
			}`,
			fieldsToCheck: []string{"directoryID", "externalID", "appID"},
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			actual, err := json.Marshal(tt.query)
			// We always expect JSON marshaling to succeed, regardless of validation
			assert.NoError(s.T(), err)

			var expectedJSON, actualJSON map[string]interface{}
			err = json.Unmarshal([]byte(tt.expected), &expectedJSON)
			assert.NoError(s.T(), err)

			err = json.Unmarshal(actual, &actualJSON)
			assert.NoError(s.T(), err)

			assert.Equal(s.T(), expectedJSON, actualJSON)

			// Validate only the fields we're testing
			validators := tt.query.GetKeyValidators()
			for _, key := range tt.fieldsToCheck {
				validator, exists := validators[key]
				assert.True(s.T(), exists, "Validator not found for key %s", key)

				value := getFieldValue(&tt.query, key)
				assert.NotNil(s.T(), value, "Value not found for key %s", key)

				isValid := validator(value)
				if tt.wantErr {
					assert.False(s.T(), isValid, "Expected validation to fail for %s with value %s", key, formatValue(value))
				} else {
					assert.True(s.T(), isValid, "Expected validation to pass for %s with value %s", key, formatValue(value))
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

// Helper function to format values for error messages
func formatValue(value interface{}) string {
	if value == nil {
		return "<nil>"
	}

	switch v := value.(type) {
	case *string:
		if v == nil {
			return "<nil>"
		}
		return *v
	case *time.Time:
		if v == nil {
			return "<nil>"
		}
		return v.String()
	case *int:
		if v == nil {
			return "<nil>"
		}
		return strconv.Itoa(*v)
	case *int32:
		if v == nil {
			return "<nil>"
		}
		return strconv.FormatInt(int64(*v), 10)
	case *int64:
		if v == nil {
			return "<nil>"
		}
		return strconv.FormatInt(*v, 10)
	case *bool:
		if v == nil {
			return "<nil>"
		}
		return strconv.FormatBool(*v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
