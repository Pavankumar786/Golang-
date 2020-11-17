Example 1: Channels
package main  
import "fmt"  
import "time"  
func worker(done chan bool) {  
fmt.Print("working...")  
time.Sleep(time.Second)  
fmt.Println("done")  
done <- true  
}  
func main() {  
done := make(chan bool, 1)  
go worker(done)  
<-done  
}  
Output:
working...done

Example 2:
package main 
import  ( 
      "fmt"
      "runtime")
func sendMessage(channelInstance chan int) {
channelInstance<-40
channelInstance<-41
}
func receiveMessage() {
}
func main() {
channelInstance:=make(chan int,3)
go func() {
channelInstance<-40
channelInstance<-41
} ()
fmt.Println("The number of channels is",runtime.NumGoroutine())
value:=<-channelInstance
fmt.Println("The value consumed from channel is:",value)
}

Output:
The number of channels is 2

Example 3:
package main 
import "fmt"
func main() { 
var mychannel chan int
fmt.Println("Value of the channel: ", mychannel) 
fmt.Printf("Type of the channel: %T ", mychannel) 
mychannel1 := make(chan int) 
fmt.Println("\nValue of the channel1: ", mychannel1) 
fmt.Printf("Type of the channel1: %T ", mychannel1) 
} 
Output:
Value of the channel:  
Type of the channel: chan int 
Value of the channel1:  0x432080
Type of the channel1: chan int 

Example 4:
package main 
import "fmt"
func myfunc(ch chan int) { 
fmt.Println(234 + <-ch) 
} 
func main() { 
    fmt.Println("start Main method") 
    // Creating a channel 
    ch := make(chan int) 
  
    go myfunc(ch) 
    ch <- 23 
    fmt.Println("End Main method") 
} 
Output:
start Main method
257
End Main method

Example 5:
package main 
import "fmt"
func myfun(mychnl chan string) { 
for v := 0; v < 4; v++ { 
mychnl <- "Hi..Pavan"
    } 
    close(mychnl) 
} 
func main() { 
c := make(chan string) 
go myfun(c) 
for { 
res, ok := <-c 
if ok == false { 
fmt.Println("Channel Close ", ok) 
brea
} 
fmt.Println("Channel Open ", res, ok) 
} 
} 
Output:

Channel Open  Hi..Pavan true
Channel Open  Hi..Pavan true
Channel Open  Hi..Pavan true
Channel Open  Hi..Pavan true
Channel Close  false


Example 6:
package main 
import "fmt"
// Main function 
func main() { 
 mychnl := make(chan string) 
 go func() { 
        mychnl <- "GFG"
        mychnl <- "gfg"
        mychnl <- "Geeks"
        mychnl <- "GeeksforGeeks"
        close(mychnl) 
    }() 
  
    // Using for loop 
    for res := range mychnl { 
        fmt.Println(res) 
    } 
} 
Output:
GFG
gfg
Geeks
GeeksforGeeks

Example 7:
package main 
import "fmt"
func main() { 
 mychnl := make(chan string, 5) 
 mychnl <- "GFG"
 mychnl <- "gfg"
 mychnl <- "Geeks"
 mychnl <- "GeeksforGeeks"
 fmt.Println("Capacity of the channel is: ", cap(mychnl)) 
} 
Output:

Capacity of the channel is:  5

//7.example
package main
import "fmt"
func main() {
    messages := make(chan string)

    go func() { messages <- "ping" }()


    msg := <-messages
    fmt.Println(msg)
}
$go run main.go
hi how r u

//8.example
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum 
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c 

	fmt.Println(x, y, x+y)
}
$go run main.go
-5 17 12

//9.example
package main 
    
import "fmt"
    
func main() { 
  
    // Creating a channel using make() function 
    ch:= make(chan int, 5) 
    fmt.Printf("\nChannel: %d", len(ch)) 
    ch <- 0 
    ch <- 1 
    ch <- 2 
    fmt.Printf("\nChannel: %d", len(ch))
} 
$go run main.go

Channel: 0
Channel: 3



