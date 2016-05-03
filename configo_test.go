package configo

import "testing"

const DATA string = `
{
    "addr": {"host": "", "port": 8080},
    "name": {"first": "Hongseok"},
    "hair_color": "black" 
}`

var c *Config
var err error

func TestNewConfigFromData(t *testing.T) {
	c, err = NewConfigFromData([]byte(DATA + ";;;"))
	if err == nil {
		t.Error("Invalid json data should be caused as failure")
	}

	c, err = NewConfigFromData([]byte(DATA))
	if err != nil {
		t.Errorf("Failed to read config:%v", err)
	}
}

func TestGet(t *testing.T) {
	port := c.GetInt("addr.port", 8000)
	if port != 8080 {
		t.Errorf("addr.port is %d, expected 8080", port)
	}

	host := c.GetString("addr.host", "test")
	if host != "" {
		t.Errorf("addr.host is [%v], expected empty string", host)
	}

	hairColor := c.GetString("hair_color", "brown")
	if hairColor != "black" {
		t.Errorf("hair_color is %v, expected black", hairColor)
	}
}
