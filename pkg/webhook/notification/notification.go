package notification

// https://goharbor.io/docs/1.10/working-with-projects/project-configuration/configure-webhooks/

type EventType string

const (
	PushImageEvent          EventType = "pushImage"
	PullImageEvent          EventType = "pullImage"
	DeleteImageEvent        EventType = "deleteImage"
	UploadChartEvent        EventType = "uploadChart"
	DownloadChartEvent      EventType = "downloadChart"
	DeleteChartEvent        EventType = "deleteChart"
	ScanImageCompletedEvent EventType = "scanningCompleted"
	ScanImageFailedEvent    EventType = "scanningFailed"
	ProjectQuotaExceed      EventType = "exceedQuota"
)

func GetEventType() map[EventType]struct{} {
	return map[EventType]struct{}{
		PushImageEvent:          {},
		PullImageEvent:          {},
		DeleteImageEvent:        {},
		UploadChartEvent:        {},
		DownloadChartEvent:      {},
		DeleteChartEvent:        {},
		ScanImageCompletedEvent: {},
		ScanImageFailedEvent:    {},
		ProjectQuotaExceed:      {},
	}
}

type Payload struct {
	EventType string  `json:"event_type"`
	Events    []Event `json:"events"`
}

type Event struct {
	Project     string `json:"project"`
	RepoName    string `json:"repo_name"`
	Tag         string `json:"tag"`
	ProjectType string `json:"project_type"`
	ImageId     string `json:"image_id"`
}
