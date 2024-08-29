package models

type ResponseWithMetadata struct {
	Data               interface{} `json:"data"`
	CurrentPage        int         `json:"current_page"`
	PageItems          int         `json:"items_count"`
	TotalCount         int         `json:"total_count"`
	TotalPages         int         `json:"total_pages"`
	PrevCursor         string      `json:"prev_cursor"`
	NextCursor         string      `json:"next_cursor"`
	RateLimitLimit     int         `json:"rate_limit_limit"`
	RateLimitRemaining int         `json:"rate_limit_remaining"`
	RateLimitReset     int         `json:"rate_limit_reset"`
	Error              error       `json:"error"`
}
