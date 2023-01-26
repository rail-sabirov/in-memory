package cache

type cache struct {
	dt map[string]interface{}
}

func New() cache {

	return cache{
		dt: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.dt[key] = value
}

func (c cache) Get(key string) interface{} {
	if _, ok := c.dt[key]; ok {
		return c.dt[key]
	}

	return nil
}

func (c *cache) Delete(key string) {
	delete(c.dt, key)
}
