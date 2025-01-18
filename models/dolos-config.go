package models

type DolosConfig struct {
	Scene      int      `json:"scene"`
	Parameters []string `json:"parameters"`
	Chash      string   `json:"chash"`
	Overdue    int      `json:"overdue"`
	DataUpload int      `json:"data_upload"`
}
