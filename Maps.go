Example 1: GO Maps

package main
import "fmt"
func main() {
var countryCapitalMap map[string]string
countryCapitalMap = make(map[string]string)
countryCapitalMap["France"] = "Paris"
 countryCapitalMap["Italy"] = "Rome"
 countryCapitalMap["Japan"] = "Tokyo"
countryCapitalMap["India"] = "New Delhi"
 for country := range countryCapitalMap {
 fmt.Println("Capital of",country,"is",countryCapitalMap[country])
 }
capital, ok := countryCapitalMap["United States"]
if(ok){
fmt.Println("Capital of United States is", capital)  
} else {
 fmt.Println("Capital of United States is not present") 
}
}
Output:
Capital of Japan is Tokyo
Capital of India is New Delhi
Capital of France is Paris
Capital of Italy is Rome
Capital of United States is not present

Example 2:
package main  
import "fmt"  
func main ()  {  
x := map[string]int{"Kate":28,"John":37, "Raj":20}  
fmt.Print(x)  
fmt.Println("\n",x["Raj"])  
}  

output:
map[John:37 Raj:20 Kate:28]
 20


Example 3:
package main  
import "fmt"  
func main ()  {  
x := map[string]int{"Pavan":24,"kumar":23, "Raj":20}  
fmt.Print(x)  
fmt.Println("\n",x["Raj"])  
}  

Output:
map[kumar:23 Raj:20 Pavan:24]
20


Example 4:Go map Insert and Update operation
package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   fmt.Println(m)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   m["Key2"] = 555  
   fmt.Println(m)  
}  
Output:
map[]
map[Key3:30 Key1:10 Key2:20]
map[Key1:10 Key2:555 Key3:30]


Example 5:

package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   fmt.Println(m)  
   m["Key1"] = 100 
   m["Key2"] = 200  
   m["Key3"] = 300  
   fmt.Println(m)  
   m["Key2"] = 555  
   fmt.Println(m)  
}  

Output:
map[]
map[Key1:100 Key2:200 Key3:300]
map[Key1:100 Key2:555 Key3:300]


Example 6:Go Map Delete operation
package main  
import "fmt"  
func main() {  
m := make(map[string]int)  
m["Key1"] = 10  
m["Key2"] = 20  
m["Key3"] = 30  
fmt.Println(m)  
delete(m, "Key3")  
fmt.Println(m)  
}

Output:
map[Key1:10 Key2:20 Key3:30]
map[Key2:20 Key1:10]


Example 7:Go Retrive Element
package main  
import "fmt"  
func main() {  
m := make(map[string]int)  
m["Key1"] = 10  
m["Key2"] = 20  
m["Key3"] = 30  
fmt.Println(m)  
fmt.Println("The value:", m["Key2"])  
}  
Output:
map[Key1:10 Key2:20 Key3:30]
The value: 20
     

Example 8:Go map struct
package main  
import "fmt"  
type Vertex struct {  
Latitude, Longitude float64  
}  
var m = map[string]Vertex{  
"Java": Vertex{     40.68433, -74.39967,   },  
"SSS-IT": Vertex{     37.42202, -122.08408,  },  
}  
func main() {  
fmt.Println(m)  
}  
Output:
map[Java:{40.68433 -74.39967} SSS-IT:{37.42202 -122.08408}]

Example 9:
package main  
import "fmt"  
func main() {  
m := make(map[string]int)  
m["Key1"] = 10  
m["Key2"] = 20  
m["Key3"] = 30  
fmt.Println(m)  
v, ok := m["Key2"]  
fmt.Println("The value:", v, "Present?", ok)  
i, j := m["Key4"]  
fmt.Println("The value:", i, "Present?", j)  
}  
Output:
map[Key1:10 Key2:20 Key3:30]
The value: 20 Present? true
The value: 0 Present? false




