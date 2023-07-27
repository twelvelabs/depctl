package core

type Dependency struct {
	Name        string `yaml:"name"`
	Install     string `yaml:"install"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
}

type Dependencies map[string][]Dependency
