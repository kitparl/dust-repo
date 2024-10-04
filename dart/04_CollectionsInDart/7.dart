
// - Create a map with name, phone keys and store some values to it. Use where to find all keys that have length 4.

void main(){
    Map<String, dynamic> map = {
        "name": "Pranshu",
        "phone": 123456
      };

    for(var key in map.keys){
      if(key.length == 4){
        print(key);
      }
    }
  }
