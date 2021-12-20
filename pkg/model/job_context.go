package model

type JobContext struct {
	Status    string                       `json:"status"`
	Container jobContainerContext          `json:"container"`
	Services  map[string]jobServiceContext `json:"services"`
}

type JobContainerContext struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

type JobServiceContext struct {
	ID      string            `json:"id"`
	Network string            `json:"network"`
	Ports   map[string]string `json:"ports"`
}
