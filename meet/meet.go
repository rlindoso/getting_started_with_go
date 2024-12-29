package meet

import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, people!!")
		time.Sleep(100 * time.Millisecond)
	}
}
