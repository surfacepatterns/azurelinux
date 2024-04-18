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
//	fmt.Println("[DEBUG] base is ", base_dir)
//	fmt.Println("[DEBUG] wd is ", wd)
	setConfigMap("PROJECT_ROOT", base_dir)

	lines, err := file.ReadLines(wd+"/"+configFile)
//	fmt.Println("[DEBUG] opened file: ", len(lines))
	if err != nil {
		return fmt.Errorf("failed to open file:\n%w", err)
	}
	for _, line := range lines {
//		fmt.Println("[DEBUG] line is", line)
		entry := strings.Split(line,":")
		if len(entry) != 2 {
//			fmt.Println("[ERROR] not a config entry", entry[0])
			continue
		}
		entry[1] = strings.Replace(entry[1], "<PROJECT_ROOT>/", base_dir, -1)
		setConfigMap(entry[0], entry[1])
//		fmt.Println("[DEBUG] entry is is", entry[0], ":",entry[1] )
		i,_ := getConfigMap(entry[0])
//		fmt.Println("[DEBUG] returnied ",i)
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
