package terminal

type Option struct {
	Image   string `json:"image,omitempty" yaml:"image,omitempty"`
	Timeout int    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

func NewTerminalOption() *Option {
	return &Option{
		Image:   "alpine:3.15",
		Timeout: 600,
	}
}
