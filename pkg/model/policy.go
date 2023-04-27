package model

type Policy struct {
	Sub string `json:"role"`
	Obj string `json:"url"`
	Act string `json:"method"`
}
