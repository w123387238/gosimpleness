package drawproperty

import (
	"bufio"
	"errors"

	"os"
	"regexp"
	"strings"
)

var rule string = "^\\s*[_A-Za-z]+[_A-Za-z0-9]*(.[_A-Za-z0-9]+)*\\s*="
var regcop *regexp.Regexp

func init() {
	regcop, _ = regexp.Compile(rule)
}
func LoadProperties(fileUri string, propertiesMap map[string]string) error {
	f, err := os.Open(fileUri)
	if err != nil {
		return errors.New("can not load file!")
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		line, prefix, readLineErr := br.ReadLine()
		if readLineErr != nil {
			break
		} else if prefix == true {
			break
		} else {
			lineStr := string(line[:])
			if strings.HasPrefix(strings.TrimSpace(lineStr), "#") {
				continue
			} else if regcop.MatchString(lineStr) {
				equalityIndex := strings.Index(lineStr, "=")
				key := strings.TrimSpace(lineStr[0:equalityIndex])
				value := strings.TrimSpace(lineStr[equalityIndex+1:])
				propertiesMap[key] = value
			}
		}
	}
	return nil
}
