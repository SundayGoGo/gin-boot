package models

type Access struct {
	BaseModel
	Ip      string `json:"ip"`
	Pro     string `json:"pro"` // 省份
	City    string `json:"city"`
	Addr    string `json:"addr"`
	RawJson string `gorm:"type:text" json:"raw_json"` // 请求返回的原始json
}
