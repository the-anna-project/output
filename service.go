package output

// ServiceConfig represents the configuration used to create a new service.
type ServiceConfig struct {
	// Dependencies.
	Channel chan Output
}

// DefaultServiceConfig provides a default configuration to create a new service
// by best effort.
func DefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		// Dependencies.
		Channel: make(chan Output, 1),
	}
}

// NewService creates a new configured service.
func NewService(config ServiceConfig) (Service, error) {
	// Dependencies.
	if config.Channel == nil {
		return nil, maskAnyf(invalidConfigError, "channel must not be empty")
	}

	newService := &service{
		// Dependencies.
		channel: config.Channel,
	}

	return newService, nil
}

type service struct {
	// Dependencies.
	channel chan Output
}

func (s *service) Channel() chan Output {
	return s.channel
}
