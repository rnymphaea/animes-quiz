package models

type Anime struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desciption"`
	Episodes uint16 `json:"episodes"`
	Genre    string `json:"genre"`
	Status   string `json:"status"`
	Type     string `json:"type"`
}
