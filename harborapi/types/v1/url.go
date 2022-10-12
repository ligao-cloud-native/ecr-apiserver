package v1

const (
	DeleteChart             = "/chartrepo/{repo}/charts/{name}"
	DeleteChartVersion      = "/chartrepo/{repo}/charts/{name}/{version}"
	DeleteChartVersionLabel = "/chartrepo/{repo}/charts/{name}/{version}/labels/{id}"
)
