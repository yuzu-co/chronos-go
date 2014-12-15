package chronos

type Container struct {
	Type    string `json:"type"`
	Image   string `json:"image"`
	Network string `json:"network"`
}

// NewContainer creates a new Container assignment
func NewContainer(image string) *Container {
	return &Container{
		Type:    "DOCKER",
		Image:   image,
		Network: "BRIDGE",
	}
}
