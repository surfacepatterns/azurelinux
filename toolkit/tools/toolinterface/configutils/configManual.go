// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package configutils

import (
	"fmt"
	"os"
	"strings"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/file"

)

var configMap map[string]string
const configFile = "configutils/config.txt"

func populateConfigFromFile() (err error) {
	configMap = make(map[string]string)
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory:\n%w", err)
	}
	base_dir := strings.Split(wd, "toolkit")[0]
	logger.Log.Debugf("base is ", base_dir)
	logger.Log.Debugf("wd is ", wd)
	SetConfigMap("PROJECT_ROOT", base_dir)

	lines, err := file.ReadLines(wd+"/"+configFile)
	logger.Log.Debugf("opened file: ", len(lines))
	if err != nil {
		return fmt.Errorf("failed to open file:\n%w", err)
	}
	for _, line := range lines {
		logger.Log.Debugf("line is", line)
		entry := strings.Split(line,":")
		if len(entry) != 2 {
			logger.Log.Errorf("not a config entry", entry[0])
			continue
		}
		entry[1] = strings.Replace(entry[1], "<PROJECT_ROOT>/", base_dir, -1)
		SetConfigMap(entry[0], entry[1])
		logger.Log.Debugf("entry is is", entry[0], ":",entry[1] )
		i,_ := GetConfigMap(entry[0])
		logger.Log.Debugf("returnied ",i)
	}
	return
}

func setConfigMap(key, val string) {
	configMap[key] = val
	return
}

func getConfigMap(key string) (val string, err error) {
	val, exists := configMap[key]
	if !exists {
		err = fmt.Errorf("key does not exist")
	}
	return
}
