package models

type ResponseWithMetadata struct {
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page"`
	PageItems   int         `json:"items_count"`
	TotalCount  int         `json:"total_count"`
	TotalPages  int         `json:"total_pages"`
	PrevCursor  string      `json:"prev_cursor"`
	NextCursor  string      `json:"next_cursor"`
	Error       error       `json:"error"`
}
