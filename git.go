package padArchiver

import (
	"fmt"
	"os/exec"
)

//TODO this is not finished
func (repo *Repo) GitUpdate(commitMsg string) error {
	_, err := exec.Command("bash", "-c", "git pull origin master").Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = exec.Command("bash", "-c", "git add .").Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = exec.Command("bash", "-c", "git commit -m '"+commitMsg+"'").Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
