// Package configo privides accessbility for JSON-formatted config file in the simplest way.
// Value can be accessed with Get... function and dot(.) seperated path string
package configo

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Config struct {
	data map[string]interface{}
}

func NewConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewConfigFromData(data)
}

func NewConfigFromData(data []byte) (*Config, error) {
	c := &Config{}

	err := json.Unmarshal(data, &c.data)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) get(path string) interface{} {
	keys := strings.Split(path, ".")

	var val interface{} = c.data
	for _, key := range keys {
		v, ok := val.(map[string]interface{})
		if !ok {
			return nil
		}

		val = v[key]
	}

	return val
}

func (c *Config) GetInt(path string, defaultValue int) int {
	val := c.get(path)

	v, ok := val.(float64)
	if !ok {
		return defaultValue
	}

	return int(v)
}

func (c *Config) GetFloat64(path string, defaultValue float64) float64 {
	val := c.get(path)

	v, ok := val.(float64)
	if !ok {
		return defaultValue
	}

	return v
}

func (c *Config) GetString(path string, defaultValue string) string {
	val := c.get(path)

	v, ok := val.(string)
	if !ok {
		return defaultValue
	}

	return v
}
