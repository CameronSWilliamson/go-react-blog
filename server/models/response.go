package models

type Response struct {
	Content interface{} `json:"content"`
}

type SampleResponse struct {
	ID     string
	Status string
}
