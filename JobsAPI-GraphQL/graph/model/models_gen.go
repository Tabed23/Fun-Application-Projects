// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type CreateJobResponse struct {
	Status int    `json:"status"`
	ID     string `json:"ID"`
	Msg    string `json:"msg"`
}

type DeleteJobResponse struct {
	Status      int    `json:"status"`
	DeleteJobID string `json:"deleteJobID"`
	Msg         string `json:"msg"`
}

type JobList struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type UpdateJobInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

type UpdateJobResponse struct {
	Status     int      `json:"status"`
	UpdatedJob *JobList `json:"updatedJob,omitempty"`
	Msg        string   `json:"msg"`
}
