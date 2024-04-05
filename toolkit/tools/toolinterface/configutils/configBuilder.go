// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package configutils

import (
	"fmt"
	"strings"
	"os"

    config "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"

)

const dirConfigFile = "configutils/directory_configs.json"
var dirConfig = config.New("dir-config")

func SetupConfig() (err error) {
	dirConfig.WithOptions(config.ParseEnv)
	dirConfig.AddDriver(json.Driver)
	err = dirConfig.LoadFilesByFormat("json", dirConfigFile)
	if err != nil {
		err = fmt.Errorf("failed to load config from file (%s):\n%w", dirConfigFile, err)
	}
	baseDir, wd, err := getBaseDir()
	if err != nil {
		return
	}
	setConfig(dirConfig, "PROJECT_ROOT", baseDir)
	replaceConfig(dirConfig, "<PROJECT_ROOT>", baseDir)
	fmt.Println("[debug] working dir is:", wd)
	fmt.Println("[debug] ************** config data:\n", dirConfig.Data())
	return
}

func replaceConfig(c *config.Config, old, new string) (err error) {
	mapData := c.Data()
	for k, v := range mapData {
		if strings.Contains(v.(string), old) {
			err = setConfig(c, k, strings.Replace(v.(string), old, new, 1))
		}
    }
	return
}

func setConfig(c *config.Config, key, val string) (err error){
	return c.Set(key, val)
}

func getConfig(c *config.Config, key string) (val string, err error) {
	exists := c.Exists(key)
	if !exists {
		return "", fmt.Errorf("failed to get value as key does not exist")
	}
	val = c.String(key)
	return
}

func GetBuildConfig(key string) (val string, err error) {
	return getConfig(dirConfig, key)
}

func SetBuildConfig(key, val string) (err error) {
	return setConfig(dirConfig, key, val)
}

func getBaseDir() (baseDir, wd string, err error){
	wd, err = os.Getwd()
	if err != nil {
		err = fmt.Errorf("failed to get working directory:\n%w", err)
		return
	}
	baseDir = strings.Split(wd, "toolkit")[0]
	return
}
