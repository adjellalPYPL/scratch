package SC_727

import (
	"fmt"
	"time"
)

func main() {
	AddNums()
}

func AddNums() int {
	var1 := 1
	var2 := 2
	fmt.Println("running add nums func")
	time.Sleep(2 * time.Second)
	return var1 + var2
}
