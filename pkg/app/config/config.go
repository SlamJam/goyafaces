package config

import "fmt"

type Config struct {
	Foo *Foo
	Baz int
	Dox []string
	Box []int
	Fox map[string]*Foo
	Gox []*Foo
	key string
}

func (c *Config) Default() {
	if c.Foo == nil {
		c.Foo = &Foo{Bar: "bar default", Qux: true}
	}

	// if c.Fox == nil {
	// 	c.Fox = map[string]*Foo{"key": {}}
	// }

	if c.Gox == nil {
		c.Gox = []*Foo{{}}
	}
}

func (c *Config) Validate() error {
	if c.Baz <= 0 {
		return fmt.Errorf("baz should be greater than zero")
	}
	return nil
}

func (c *Config) Expand() error {
	c.key = "value written by Expand()"

	return nil
}

type Foo struct {
	Bar string
	Qux bool
}

func (c *Foo) Default() {
	if c.Bar == "" {
		c.Bar = "default value"
	}
}
