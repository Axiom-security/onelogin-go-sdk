package models

type AuthServerQuery struct {
	Name string `json:"name,omitempty"`
	BaseQueryRequest
}

type AuthServer struct {
	ID            *int32                   `json:"id,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	Configuration *AuthServerConfiguration `json:"configuration,omitempty"`
}

type AuthServerConfiguration struct {
	ResourceIdentifier            *string  `json:"resource_identifier,omitempty"`
	Audiences                     []string `json:"audiences,omitempty"`
	AccessTokenExpirationMinutes  *int32   `json:"access_token_expiration_minutes,omitempty"`
	RefreshTokenExpirationMinutes *int32   `json:"refresh_token_expiration_minutes,omitempty"`
}

type ClientAppsQuery struct {
	ID int `json:"id,omitempty"`
}

type ClientApp struct {
	ID           *int32  `json:"app_id,omitempty"`
	AuthServerID *int32  `json:"auth_server_id,omitempty"`
	APIAuthID    *int32  `json:"api_auth_id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Scopes       []Scope `json:"scopes,omitempty"`
	ScopeIDs     []int32 `json:"scope_ids,omitempty"`
}

type ScopesQuery struct {
	ID int `json:"id,omitempty"`
}

type Scope struct {
	ID           *int32  `json:"id,omitempty"`
	AuthServerID *int32  `json:"auth_server_id,omitempty"`
	Value        *string `json:"value,omitempty"`
	Description  *string `json:"description,omitempty"`
}

type AccessTokenClaimsQuery struct {
	ID int `json:"id,omitempty"`
}

type AccessTokenClaim struct {
	ID                       *int32   `json:"id,omitempty"`
	AuthServerID             *int32   `json:"auth_server_id,omitempty"`
	Label                    *string  `json:"label,omitempty"`
	UserAttributeMappings    *string  `json:"user_attribute_mappings,omitempty"`
	UserAttributeMacros      *string  `json:"user_attribute_macros,omitempty"`
	AttributeTransformations *string  `json:"attribute_transformations,omitempty"`
	SkipIfBlank              *bool    `json:"skip_if_blank,omitempty"`
	Values                   []string `json:"values,omitempty"`
	DefaultValues            *string  `json:"default_values,omitempty"`
	ProvisionedEntitlements  *bool    `json:"provisioned_entitlements,omitempty"`
}

func (q *AuthServerQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"name":   validateString,
		"limit":  validateString,
		"page":   validateString,
		"cursor": validateString,
	}
}

func (q *ClientAppsQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"auth_server_id": validateString,
	}
}

func (q *ScopesQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"auth_server_id": validateString,
	}
}

func (q *AccessTokenClaimsQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"auth_server_id": validateString,
	}
}
