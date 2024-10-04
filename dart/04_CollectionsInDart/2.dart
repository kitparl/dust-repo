
// - Create a set of fruits and print all fruits using loop.

void main(){
    Set<String> fruits = {"Apple", "Banana", "Orange"};
    for(String fruit in fruits){
      print(fruit);
    }

    fruits.forEach((fruit) => print(fruit));
  }
