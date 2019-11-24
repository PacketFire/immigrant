package core

type Revision struct {
	Revision string   `yaml:"revision"`
	Migrate  []string `yaml:"migrate"`
	Rollback []string `yaml:"rollback"`
}
