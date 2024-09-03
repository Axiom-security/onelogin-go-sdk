package models

type Group struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Reference *string `json:"reference"`
}

type GroupQuery struct {
	BaseQueryRequest
}

func (p *GroupQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":  validateString,
		"page":   validateString,
		"cursor": validateString,
	}
}
