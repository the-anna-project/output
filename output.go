// Package output implements primitives to distribute output calculated by the
// neural network.
package output

// Config represents the configuration used to create a new output.
type Config struct {
	// Settings.
	Text string
}

// DefaultConfig provides a default configuration to create a new output by best
// effort.
func DefaultConfig() Config {
	return Config{
		// Settings.
		Text: "",
	}
}

// New creates a new configured output.
func New(config Config) (Output, error) {
	// Settings.
	if config.Text == "" {
		return nil, maskAnyf(invalidConfigError, "text must not be empty")
	}

	newOutput := &output{
		// Settings.
		text: config.Text,
	}

	return newOutput, nil
}

type output struct {
	// Settings.
	text string
}

func (e *output) Text() string {
	return e.text
}
