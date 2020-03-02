package bst_web_models

type DdrSongIds struct {
	Ids []string `json:"ids"`
}

type DdrSongIdWithJacket struct {
	Id string `json:"id"`
	Jacket string `json:"jacket"`
}