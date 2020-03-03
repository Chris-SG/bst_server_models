package bst_api_models

type Status struct {
	Status string `json:"status"`
	Message string `json:"message,omitempty"`
}

type ApiStatus struct {
	Api string `json:"api"`
	EaGate string `json:"gate"`
	Db string `json:"db"`
}

type Ordering struct {
	OrderBy []string `json:"order_by"`
}