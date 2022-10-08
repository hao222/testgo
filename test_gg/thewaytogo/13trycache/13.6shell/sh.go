// exec.go
package main

import (
	"fmt"
	"log"
	"os/exec"
)

//func main() {
//	// 1) os.StartProcess //
//	/*********************/
//	/* Linux: */
//	env := os.Environ()
//	procAttr := &os.ProcAttr{
//		Env: env,
//		Files: []*os.File{
//			os.Stdin,
//			os.Stdout,
//			os.Stderr,
//		},
//	}
//	// 1st example: list files
//	pid, err := os.StartProcess("/bin", []string{"ls", "-l"}, procAttr)
//	if err != nil {
//		fmt.Printf("Error %v starting process!", err) //
//		os.Exit(1)
//	}
//	fmt.Printf("The process id is %v", pid)
//}

//func main() {
//	cmd := exec.Command("ls")
//	cmd.Dir = "/Users/wuhao/mix/"
//	var out bytes.Buffer
//	var stderr bytes.Buffer
//	cmd.Stdout = &out
//	cmd.Stderr = &stderr
//	err := cmd.Run()
//	if err != nil {
//		fmt.Println(fmt.Sprint(err) + ":" + stderr.String())
//		return
//	}
//	fmt.Println("Result:" + out.String())
//}

func main() {

	out, err := exec.Command("ls", "-al").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
