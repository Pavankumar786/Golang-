Example 1:
package main  
import "encoding/json"  
import "fmt"  
  func main() {  
bolType, _ := json.Marshal(false) //boolean Value  
fmt.Println(string(bolType))  
intType, _ := json.Marshal(10) // integer value  
fmt.Println(string(intType))  
fltType, _ := json.Marshal(3.14) //float value  
fmt.Println(string(fltType))  
strType, _ := json.Marshal("JavaTpoint") // string value  
fmt.Println(string(strType))  
slcA := []string{"sun", "moon", "star"} //slice value  
slcB, _ := json.Marshal(slcA)  
fmt.Println(string(slcB))  
mapA := map[string]int{"sun": 1, "moon": 2} //map value  
 mapB, _ := json.Marshal(mapA)  
fmt.Println(string(mapB))  
}  
Output:
false
10
3.14
"JavaTpoint"
["sun","moon","star"]
{"moon":2,"sun":1

Example 2: (User Defined Data Type)
package main  
  import ("encoding/json"  
           "fmt"  
           "os")  
  
type Response1 struct {  
    Position   int  
    Planet []string  
}  
type Response2 struct {  
    Position   int      'json:"position"'  
    Planet []string 'json:"planet"'  
}  
func main()  {  
    res1A := &Response1{  
        Position:   1,  
        Planet: []string{"mercury", "venus", "earth"}}  
    res1B, _ := json.Marshal(res1A)  
    fmt.Println(string(res1B))  
  
    res2D := &Response2{  
        Position:   1,  
        Planet: []string{"mercury", "venus", "earth"}}  
    res2B, _ := json.Marshal(res2D)  
    fmt.Println(string(res2B))  
  
  
    byt := []byte('{"pi":6.13,"place":["New York","New Delhi"]}`)  
    var dat map[string]interface{}  
    if err := json.Unmarshal(byt, &dat); err != nil {  
        panic(err)  
    }  
    fmt.Println(dat)  
    num := dat["pi"].(float64)  
    fmt.Println(num)  
    strs := dat["place"].([]interface{})  
    str1 := strs[0].(string)  
    fmt.Println(str1)  
    str := `{"Position": 1, "Planet": ["mercury", "venus"]}`  
    res := Response2{}  
    json.Unmarshal([]byte(str), &res)  
    fmt.Println(res)  
    fmt.Println(res.Planet[1])  
    enc := json.NewEncoder(os.Stdout)  
    d := map[string]string{"1":"mercury" , "2": "venus"}  
    enc.Encode(d)  
  
}  
Output:
{"Position":1,"Planet":["mercury","venus","earth"]}
{"position":1,"planet":["mercury","venus","earth"]}
map[pi:6.13 place:[New York New Delhi]]
6.13
New York
{1 [mercury venus]}
venus
{"1":"mercury","2":"venus"}


Example 3:
package main
import (
    "encoding/json"
    "fmt"
)
type Book struct {
Title  string `json:"title"`
Author Author `json:"author"`
}

type Author struct {
 Sales     int  `json:"book_sales"`
 Age       int  `json:"age"`
 Developer bool `json:"is_developer"`
}
func main() {
author := Author{Sales: 3, Age: 25, Developer: true}
book := Book{Title: "Learning Concurrency in Python", Author: author}
byteArray, err := json.Marshal(book)
if err != nil {
 fmt.Println(err)
 }
fmt.Println(string(byteArray))
}

Output:
{"title":"Learning Concurrency in Python","author":{"book_sales":3,"age":25,"is_developer":true}}