package chronos

type Container struct {
	Type    string   `json:"type"`
	Image   string   `json:"image"`
	Network string   `json:"network"`
	Volumes []Volume `json:"volumes"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Volume struct {
	HostPath      string `json:"hostPath"`
	ContainerPath string `json:"containerPath"`
	Mode          string `json:"mode"`
}

// NewContainer creates a new Container assignment
func NewContainer(image string) *Container {
	return &Container{
		Type:    "DOCKER",
		Image:   image,
		Network: "BRIDGE",
	}
}
