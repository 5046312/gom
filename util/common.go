package util

import (
	"fmt"
	"os"
	"os/exec"
)

func Icon() {
	gom := `
   ######    #######  ##     ## 
  ##    ##  ##     ## ###   ### 
  ##        ##     ## #### #### 
  ##   #### ##     ## ## ### ## 
  ##    ##  ##     ## ##     ## 
   ######    #######  ##     ## 
`
	fmt.Println(gom)
}

func Command(args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Path Exist
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
