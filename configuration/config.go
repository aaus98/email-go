package configuration

//Config server port email and password
type Config struct {
	Server   string
	Port     int
	Email    string
	Password string
}

func (c *Config) Read() {
	c.Server = "smtp.gmail.com"
	c.Port = 587
	c.Email = ""
	c.Password = ""
}
