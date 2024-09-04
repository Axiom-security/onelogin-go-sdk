package models

type APICredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Subdomain    string `json:"subdomain"`
}

type ResponseMetadata struct {
	CurrentPage        int    `json:"current_page"`
	PageItems          int    `json:"items_count"`
	TotalCount         int    `json:"total_count"`
	TotalPages         int    `json:"total_pages"`
	PrevCursor         string `json:"prev_cursor"`
	NextCursor         string `json:"next_cursor"`
	RateLimitLimit     int    `json:"rate_limit_limit"`
	RateLimitRemaining int    `json:"rate_limit_remaining"`
	RateLimitReset     int    `json:"rate_limit_reset"`
}

type ResponseWithMetadata struct {
	Data     interface{}      `json:"data"`
	Metadata ResponseMetadata `json:"metadata"`
	Error    error            `json:"error"`
}

type BaseQueryRequest struct {
	Limit  string `json:"limit,omitempty"`
	Page   string `json:"page,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

func (b *BaseQueryRequest) SetLimit(limit string) {
	b.Limit = limit
}

func (b *BaseQueryRequest) SetPage(page string) {
	b.Page = page
}

func (b *BaseQueryRequest) SetCursor(cursor string) {
	b.Cursor = cursor
}
