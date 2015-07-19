package chronos

type Job struct {
	Name          string `json:"name"`
	Command       string `json:"command"`
	Shell         bool   `json:"shell,omitempty"`
	Epsilon       string `json:"epsilon"`
	Executor      string `json:"executor"`
	ExecutorFlags string `json:"executorFlags"`
	Retries       int    `json:"retries,omitempty"`
	Owner         string `json:"owner"`
	Async         bool   `json:"async,omitempty"`
	Successs      int    `json:"successCount,omitempty"`
	Errors        int    `json:"errorCount,omitempty"`
	//LastSuccess   Time     `json:"lastSuccess,omitempty"`
	//LastError     Time     `json:"lastError,omitempty"`
	CPUs     float32  `json:"cpus"`
	Disk     float32  `json:"disk"`
	Memory   float32  `json:"mem"`
	Disabled bool     `json:"disabled,omitempty"`
	URIs     []string `json:"uris"`
	//Environment map[string]string `json:"environmentVariables"`
	Environment      []EnvVar   `json:"environmentVariables"`
	RunAsUser        string     `json:"runAsUser"`
	Container        *Container `json:"container,omitempty"`
	Schedule         string     `json:"schedule"`
	scheduleTimeZone string     `json:"scheduleTimeZone"`
}

// NewJob creates a new Job assignment
func NewJob() Job {
	return Job{}
}

// List registered jobs in Chronos
func ListJobs() ([]Job, error) {
	var jobs []Job

	err := Get("/scheduler/jobs", &jobs)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// Get a job by its name
func GetJob(name string) (Job, error) {
	jobs, err := ListJobs()
	if err != nil {
		return Job{}, err
	}

	for _, job := range jobs {
		if job.Name == name {
			return job, nil
		}
	}

	return Job{}, nil
}

// Delete Job
func (j Job) Delete() error {
	return Delete("/scheduler/job/" + j.Name)
}

// Kill Jobs Tasks
func (j Job) KillAll() error {
	return Delete("/scheduler/task/kill/" + j.Name)
}

// Run a Job
func (j Job) Run() error {
	return Put("/scheduler/job/" + j.Name)
}

// Create Job
func (j *Job) Create() error {
	return Post("/scheduler/iso8601", j)
}
