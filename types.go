package main

import "time"

// Hold information from task
type Task struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	ProjectID  int64     `json:"projectID"`
	AssignedTo int64     `json:"assignedTo"`
	CreatedAt  time.Time `json:"createdAt"`
}
