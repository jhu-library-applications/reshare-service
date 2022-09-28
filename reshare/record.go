package reshare

type Records struct {
	InnerRecord []Record `json:"records"`
}

type Record struct {
	Id            string   `json:"id"`
	Title         string   `json:"title"`
	LendingStatus []string `json:"lendingStatus"`
}
