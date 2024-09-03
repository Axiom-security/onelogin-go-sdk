package models

type Group struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Reference *string `json:"reference"`
}

type GroupQuery struct {
	// Group object uses an old cursor field name
	Cursor string `json:"after_cursor,omitempty"`
	BaseQueryRequest
}

func (p *GroupQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":        validateString,
		"page":         validateString,
		"after_cursor": validateString,
	}
}

func (p *GroupQuery) SetCursor(cursor string) {
	p.Cursor = cursor
}
