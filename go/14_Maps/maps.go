package main

import "fmt"

// Maps are used to store data values in key:value pairs.
// Each element in a map is a key:value pair.
// A map is an unordered and changeable collection that does not allow duplicates.
// The length of a map is the number of its elements. You can find it using the len() function.
// The default value of a map is nil.
// Maps hold references to an underlying hash table.
// Go has multiple ways for creating maps.



func main() {

  // Create Maps Using var and :=
  var map1 = map[string]string{"name": "Pranshu", "age": "21"}
  map2 := map[string]string{"name": "Jai", "age": "21"}

  fmt.Println(map1)
  fmt.Println(map2)

  // Create Maps Using make()Function:
  var map3 = make(map[string]string)
  map3["name"] = "Pranshu"
  map3["age"] = "21"

  map4 := make(map[string]string)
  map4["name"] = "Jai"
  map4["age"] = "21"

  fmt.Println(map3)
  fmt.Println(map4)

  // Create an Empty Map
  // There are two ways to create an empty map
 var a = make(map[string]string)
 fmt.Println(a)
 var b map[string]string
 fmt.Println(b)

  fmt.Println(a == nil) // false
  fmt.Println(b == nil) // true

  /*

The map key can be of any data type for which the equality operator (==) is defined. These include:

Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)

Invalid key types are:

Slices
Maps
Functions

These types are invalid because the equality operator (==) is not defined for them.
  */

// Allowed Value Types
// The map values can be any type.

// Access Map Elements
// You can access map elements by:

var car = make(map[string]string)
car["brand"] = "Nano"
car["model"] = "Mustang"
car["year"] = "1969"

fmt.Println(car["brand"])

// Update and Add Map Elements
car["year"] = "1970"
car["color"] = "red"
fmt.Println(car["year"])

// Remove Element from Map
delete(car, "year")
fmt.Println(car)


// Check For Specific Elements in a Map
 val1, ok1 := a["brand"]
 _, ok2 := a["model"]
 fmt.Println(val1, ok1)
 fmt.Println(ok2)

 // Maps Are References
 // Maps are references to hash tables.
// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

  var carModel = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  model := carModel

  fmt.Println(carModel)
  fmt.Println(model)

  model["wheel"] = "4"

  fmt.Println(carModel)
  fmt.Println(model)

  // Iterate Over Maps
// we can use range to iterate over maps.


  example := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  for k, v := range example {
    fmt.Println(k, v)
  }


}
