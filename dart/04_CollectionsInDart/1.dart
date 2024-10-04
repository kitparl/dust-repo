
// - Create a list of names and print all names using list.

void main(){
    List<String> names = ["Pranshu", "Singh", "Bisht"];
    for(int i=0; i<names.length; i++){
      print(names[i]);
    }
    names.forEach((name) => print(name));
  }
