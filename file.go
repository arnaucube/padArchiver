package padArchiver

import (
	"io/ioutil"
)

func AddLineToFile(path string, line string) error {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(fileBytes)
	r := line + "\n\n"
	r = r + content

	err = ioutil.WriteFile(path, []byte(r), 0644)
	if err != nil {
		return err
	}
	return nil
}
