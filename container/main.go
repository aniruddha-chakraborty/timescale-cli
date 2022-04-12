package container

type Container struct {
	Services map[string]interface{}
}

func (c *Container) Set(name string, class interface{}) {
	c.Services[name] = class
}

func (c *Container) Get(name string) interface{} {
	return c.Services[name]
}

func (c *Container) Delete(name string) {
	delete(c.Services, name)
}