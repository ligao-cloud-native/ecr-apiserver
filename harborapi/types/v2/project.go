package v2

type Project struct {
	ProjectName string `json:"project_name"`
}

type ProjectResponse struct {
	UpdateTime string `json:"update_time"`
	Name       string `json:"name"`
	Deleted    bool   `json:"deleted"`
}
