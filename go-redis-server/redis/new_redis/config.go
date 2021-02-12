package new_redis

type Config struct {
	proto   string
	host    string
	port    int
}

func DefaultConfig() *Config {
	return &Config{
		proto:   "tcp",
		host:    "127.0.0.1",
		port:    6389,
	}
}

func (c *Config) Port(p int) *Config {
	c.port = p
	return c
}

func (c *Config) Host(h string) *Config {
	c.host = h
	return c
}

func (c *Config) Proto(p string) *Config {
	c.proto = p
	return c
}
