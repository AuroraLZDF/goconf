/**
 * Read the configuration file
 *
 * @copyright           (C) 2020  AuroraLZDF
 * @last modify          2020-02-24
 * @website		https://github.com/AuroraLZDF/goconf
 *
 */

package goconf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	filePath string                       //your ini file path directory+file
	ConfList map[string]map[string]string //configuration information slice
}

// Load configuration file
func InitConfig(filePath string) *Config {
	cfg := new(Config)
	cfg.filePath = filePath
	cfg.readList()

	return cfg
}

// To obtain corresponding value of the key values
func (c *Config) GetValue(section, name string) string {
	fmt.Println("confList: ", c.ConfList)
	_, ok := c.ConfList[section][name]
	if ok {
		return c.ConfList[section][name]
	} else {
		return ""
	}
}

// Set the corresponding value of the key value, if not add, if there is a key change
func (c *Config) SetValue(section, key, value string) {
	_, ok := c.ConfList[section]
	if ok {
		c.ConfList[section][key] = value
	} else {
		c.ConfList[section] = make(map[string]string)
		c.ConfList[section][key] = value
	}
}

// Delete the corresponding key values
func (c *Config) DeleteValue(section, name string) bool {
	_, ok := c.ConfList[section][name]
	if ok {
		delete(c.ConfList[section], name)
		return true
	}

	return false
}

// List all the configuration file
func (c *Config) readList() map[string]map[string]string {
	file, err := os.Open(c.filePath)
	if err != nil {
		CheckErr(err)
	}

	defer file.Close()

	c.ConfList = make(map[string]map[string]string)
	var section string
	var sectionMap map[string]string

	isFirstSection := true
	buf := bufio.NewReader(file)

	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}

		switch {
		case len(line) == 0:
		case string(line[0]) == ";": //增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']':
			if !isFirstSection {
				c.ConfList[section] = sectionMap
			} else {
				isFirstSection = false
			}
			section = strings.TrimSpace(line[1 : len(line)-1])
			sectionMap = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1 : ])
			sectionMap[strings.TrimSpace(line[0:i])] = value
		}
	}

	c.ConfList[section] = sectionMap
	return c.ConfList
}

// List all the configuration file
func (c *Config) GetAllSection() map[string]map[string]string {
	return c.ConfList
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Not found this error"
}

// Ban repeated appended to the slice method
func (c *Config) uniqueAppend(conf string) bool {
	for _, v := range c.ConfList {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}
