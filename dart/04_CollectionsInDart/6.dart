
// - Create a map with name, address, age, country keys and store values to it. Update country name to other country and print all keys and values.

void main(){
    Map<String, dynamic> map = {
        "name": "Pranshu",
        "address": "Blr",
        "age": 12,
        "country": "India"
      };

      map["country"] = "Other";
      print(map);

      // iteration over map
      for(var key in map.keys){
        print("$key: ${map[key]}");
      }
  }
