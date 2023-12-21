package entity

type News struct {
	Uuid     uint32 `json:"uuid"`
	Header   string `json:"header"`
	Content  string `json:"content"`
	ImgPath  string `json:"imgPath"`
	DateTime string `json:"dateTime"`
}
