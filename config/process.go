package configapp

import (
	"fmt"
	"strconv"
)

type ConfigApp struct {
	params map[string]ConfigParam
	values map[string]*string
}

type ConfigParam struct {
	Name        string
	ShortName   string
	Description string
	Default     string
	IsBool      bool
}

func (c *ConfigApp) AddParam(p ConfigParam) {
	if c.params == nil {
		c.params = make(map[string]ConfigParam)
		c.values = map[string]*string{}
	}
	c.params[p.Name] = p
	c.values[p.Name] = &p.Default
}

func (c *ConfigApp) GetInt(name string) (int64, error) {
	stringValue := *c.values[name]
	if stringValue == "" {
		return 0, nil
	}
	intvalue, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, fmt.Errorf("GetInt convert %s to int: %w", stringValue, err)
	}

	return int64(intvalue), nil
}

func (c *ConfigApp) GetBool(name string) (bool, error) {
	stringValue := *c.values[name]
	if stringValue == "" {
		return false, nil
	}
	boolValue, err := strconv.ParseBool(stringValue)
	if err != nil {
		return false, fmt.Errorf("GetBool convert %s to int: %w", stringValue, err)
	}

	return boolValue, nil
}

func (c *ConfigApp) GetString(name string) string {
	return *c.values[name]
}

func GetConfig(c ConfigApp) ConfigApp {
	if len(c.params) > 0 {
		c.values = parseParams(
			parseEnv(
				c.values,
				c.params,
			),
			c.params,
		)
	}
	return c
}
