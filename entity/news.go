package entity

type News struct {
	Uuid    uint32 `json:"uuid"`
	Header  string `json:"header"`
	Content string `json:"content"`
}
