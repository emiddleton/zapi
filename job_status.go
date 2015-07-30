package zapi

import ()

type Result struct {
	Id      int64
	Title   string
	Action  string
	Errors  string
	Success bool
	Status  string
}

type JobStatus struct {
	id       string   // yes 	no 	Automatically assigned as job gets enqueued
	url      string   // yes 	no 	The URL to poll for status updates
	total    int64    // yes 	no 	The total number of tasks this job is batching through
	progress int64    // yes 	no 	Number of tasks that have already been completed
	status   string   // yes 	no 	The current status, "working", "failed", "completed", "killed"
	message  string   // yes 	no 	Message from the job worker, if any
	results  []Result // yes 	no 	Result data from processed tasks
}
