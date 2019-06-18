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

// func Usage(args ...string) {
// 	if len(args) == 0 {
// 		fmt.Println("0 Should Show All Usage")
// 	} else if len(args) != 1 {
// 		fmt.Println("!=1 All Usage")
// 	} else {
// 		cmd, exist := command.List[args[0]]
// 		fmt.Println(args, cmd, exist)
// 		if exist {
// 			fmt.Println(cmd.Usage)
// 		} else {
// 			fmt.Println("Don't Know")
// 		}
// 	}
// }
