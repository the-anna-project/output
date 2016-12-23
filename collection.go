package output

// CollectionConfig represents the configuration used to create a new
// collection.
type CollectionConfig struct {
	// Dependencies.
	Channel chan Output
}

// DefaultCollectionConfig provides a default configuration to create a new
// collection by best effort.
func DefaultCollectionConfig() CollectionConfig {
	return CollectionConfig{
		// Dependencies.
		Channel: make(chan Output, 1),
	}
}

// NewCollection creates a new configured collection.
func NewCollection(config CollectionConfig) (*Collection, error) {
	// Dependencies.
	if config.Channel == nil {
		return nil, maskAnyf(invalidConfigError, "channel must not be empty")
	}

	var err error

	var textService Service
	{
		textConfig := DefaultServiceConfig()
		textConfig.Channel = config.Channel
		textService, err = NewService(textConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	newCollection := &Collection{
		// Public.
		Text: textService,
	}

	return newCollection, nil
}

type Collection struct {
	// Public.
	Text Service
}
