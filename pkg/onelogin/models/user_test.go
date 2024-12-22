package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserAppTestSuite struct {
	suite.Suite
}

func TestUserAppSuite(t *testing.T) {
	suite.Run(t, new(UserAppTestSuite))
}

func (s *UserAppTestSuite) TestUserApp_UnmarshalJSON() {
	tests := []struct {
		name    string
		json    string
		want    UserApp
		wantErr bool
	}{
		{
			name: "successfully unmarshal large login_id",
			json: `{
				"id": 123,
				"login_id": 2147531300,
				"icon_url": "https://example.com/icon.png"
			}`,
			want: UserApp{
				ID:      intPtr(123),
				LoginID: int64Ptr(2147531300),
				IconURL: strPtr("https://example.com/icon.png"),
			},
			wantErr: false,
		},
		{
			name: "successfully unmarshal max int64",
			json: `{
				"login_id": 9223372036854775807
			}`,
			want: UserApp{
				LoginID: int64Ptr(9223372036854775807),
			},
			wantErr: false,
		},
		{
			name: "successfully unmarshal zero login_id",
			json: `{
				"login_id": 0
			}`,
			want: UserApp{
				LoginID: int64Ptr(0),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			var got UserApp
			err := json.Unmarshal([]byte(tt.json), &got)

			if tt.wantErr {
				assert.Error(s.T(), err)
				return
			}

			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tt.want, got)
		})
	}
}

// Helper functions for creating pointers
func int64Ptr(i int64) *int64 {
	return &i
}

func intPtr(i int32) *int32 {
	return &i
}

func strPtr(s string) *string {
	return &s
}
