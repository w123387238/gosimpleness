package testter

import (
	"main/go/common/utils/configable"
	logger "main/go/common/utils/log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	propertiesMap := make(map[string]string)
	propertyErr := drawproperty.LoadProperties("../resources/configuration_testfile.properties", propertiesMap)
	if propertyErr != nil {
		t.Log(propertyErr)
		logger.Error("some bad thing!")
	}
	for k, v := range propertiesMap {
		t.Log("key:" + k + ",value:" + v)
		logger.Info("key:" + k + ",value:" + v)
	}
	logger.Flush()
}
