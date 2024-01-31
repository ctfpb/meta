package meta

type Challenge struct {
	Name  string   `yaml:"name"`
	Refer string   `yaml:"refer"`
	Tags  []string `yaml:"tags"`
}

func (c *Challenge) parseFormat() *Challenge {
	return c
}
