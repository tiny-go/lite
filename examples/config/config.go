package config

import "time"

// Config values.
type Config struct {
	Secret      string        `default:"secret"`
	TokLifeTime time.Duration `default:"1h50m"`
}
