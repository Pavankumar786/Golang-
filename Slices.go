1.program
package main
import "fmt"
func main() {
   var numbers = make([]int,3,5)
   printSlice(numbers)
}
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
output:
len=3 cap=5 slice=[0 0 0]

2.program:
package main

import "fmt"
func main() {
   var numbers []int
   printSlice(numbers)
   
   if(numbers == nil){
      fmt.Printf("slice is nil")
   }
}
func printSlice(x []int){
   fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
}
output:
len = 0 cap = 0 slice = []
slice is nil

3.program:
package main
import "fmt"
func main() {
   
   numbers := []int{0,1,2,3,4,5,6,7,8}   
   printSlice(numbers)
   fmt.Println("numbers ==", numbers)
   
   fmt.Println("numbers[1:4] ==", numbers[1:4])
   
   fmt.Println("numbers[:3] ==", numbers[:3])
  
   fmt.Println("numbers[4:] ==", numbers[4:])
   
   numbers1 := make([]int,0,5)
   printSlice(numbers1)
   number2 := numbers[:2]
   printSlice(number2)
   number3 := numbers[2:5]
   printSlice(number3)
   
}
func printSlice(x []int){
   fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
}
output:
len = 9 cap = 9 slice = [0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
numbers[1:4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
len = 0 cap = 5 slice = []
len = 2 cap = 9 slice = [0 1]
len = 3 cap = 7 slice = [2 3 4]

4.program:
package main
import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)
   numbers = append(numbers, 0)
   printSlice(numbers)
   numbers = append(numbers, 1)
   printSlice(numbers)
   numbers = append(numbers, 2,3,4)
   printSlice(numbers) 
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)  
   printSlice(numbers1)   
}
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
output:
len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 0 0 0 0]

5.program:
ackage main  
import "fmt" 
func main() { 
     
    arr := [7]string{"This", "is", "the", "haritha", 
                         "of", "Go", "language"} 
   
    fmt.Println("Array:", arr) 
    
    myslice := arr[1:6] 
  
    fmt.Println("Slice:", myslice) 
    
    fmt.Printf("Length of the slice: %d", len(myslice)) 
    
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice)) 
} 
output:
Array: [This is the haritha of Go language]
Slice: [is the haritha of Go]
Length of the slice: 5
Capacity of the slice: 6

6.program:
package main   
import "fmt"
  
func main() 
{   
    var my_slice_1 = []string{"hari", "for", "hari"} 
    fmt.Println("My Slice 1:", my_slice_1)  
    my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45} 
    fmt.Println("My Slice 2:", my_slice_2) 
} 
output:
My Slice 1: [hari for hari]
My Slice 2: [12 45 67 56 43 34 45]

7.program:
package main  
import "fmt" 
func main() { 
  
   
    arr := [4]string{"Hari", "for", "Hari", "GFG"} 
    var my_slice_1 = arr[1:2] 
    my_slice_2 := arr[0:] 
    my_slice_3 := arr[:2] 
    my_slice_4 := arr[:] 
    fmt.Println("My Array: ", arr) 
    fmt.Println("My Slice 1: ", my_slice_1) 
    fmt.Println("My Slice 2: ", my_slice_2) 
    fmt.Println("My Slice 3: ", my_slice_3) 
    fmt.Println("My Slice 4: ", my_slice_4) 
} 
output:
My Array:  [Hari for Hari GFG]
My Slice 1:  [for]
My Slice 2:  [Hari for Hari GFG]
My Slice 3:  [Hari for]
My Slice 4:  [Hari for Hari GFG]

8.program:
package main  
import "fmt"
  
func main() { 
  
   
    myslice := []string{"This", "is", "the", 
        "haritha", "of", "Go", "language"} 
 
    for _, ele := range myslice { 
        fmt.Printf("Element = %s\n", ele) 
    } 
} 
output:
Element = This
Element = is
Element = the
Element = haritha
Element = of
Element = Go
Element = language

9.program:
package main 
  
import "fmt"
  
func main() { 
  
    var myslice []string 
    fmt.Printf("Length = %d\n", len(myslice)) 
    fmt.Printf("Capacity = %d ", cap(myslice)) 
  
} 
output:
Length = 0
Capacity = 0 

10.program:
package main  
import "fmt" 
func main() { 
  
    
    arr := [6]int{55, 66, 77, 88, 99, 22} 
    slc := arr[0:4]  
    fmt.Println("Original_Array: ", arr) 
    fmt.Println("Original_Slice: ", slc) 
    slc[0] = 100 
    slc[1] = 1000 
    slc[2] = 1000 
    fmt.Println("\nNew_Array: ", arr) 
    fmt.Println("New_Slice: ", slc) 
} 
output:
Original_Array:  [55 66 77 88 99 22]
Original_Slice:  [55 66 77 88]

New_Array:  [100 1000 1000 88 99 22]
New_Slice:  [100 1000 1000 88]

11.program:
package main  
import "fmt"
  
func main() { 
  
    
    s1 := []int{12, 34, 56} 
    var s2 []int   
    /*s3:= []int{23, 45, 66} 
      fmt.Println(s1==s3) 
    */  
    fmt.Println(s1 == nil) 
    fmt.Println(s2 == nil) 
} 
output:
false
true

12.program:
package main 
  
import ( 
    "fmt"
    "sort"
) 
  
func main() { 
     
    slc1 := []string{"Python", "Java", "C#", "Go", "Hari"} 
    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}   
    fmt.Println("Before sorting:") 
    fmt.Println("Slice 1: ", slc1) 
    fmt.Println("Slice 2: ", slc2) 
    sort.Strings(slc1) 
    sort.Ints(slc2)  
    fmt.Println("\nAfter sorting:") 
    fmt.Println("Slice 1: ", slc1) 
    fmt.Println("Slice 2: ", slc2) 
  
} 
output:
Before sorting:
Slice 1:  [Python Java C# Go Ruby]
Slice 2:  [45 67 23 90 33 21 56 78 89]
After sorting:
Slice 1:  [C# Go Java Python Ruby]
Slice 2:  [21 23 33 45 56 67 78 89 90]