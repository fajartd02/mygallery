package entity

type MemoryListRequest struct {
	Filter struct {
		By      string `json:"by,omitempty"`
		Keyword string `json:"keyword,omitempty"`
	} `json:"filter,omitempty"`
	Sort struct {
		By    string `json:"by,omitempty"`
		Order string `json:"order,omitempty"`
	} `json:"sort,omitempty"`
}
