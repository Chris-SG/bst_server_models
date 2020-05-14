package bst_models

type ApiStatus struct {
	Api string `json:"api"`
	EaGate string `json:"gate"`
	Db string `json:"db"`
}

type Ordering struct {
	OrderBy []string `json:"order_by"`
}

type UserCache struct {
	Id int `json:"id"`
	Nickname string `json:"nickname"`
	Public bool `json:"public"`
}