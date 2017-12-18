package codecov_compare

type Config struct {
	data map[string]map[string]string
}

func NewConfig() () {
	_ = &Config{
		data: make(map[string]map[string]string),
	}
}

func (c *Config) parse(fname string) (err error) {
	return nil
}