
// - Add your 7 friend names to the list. Use where to find a name that starts with alphabet a.

void main(){
    List<String> friends = ["Pranshu", "Singh", "Bisht", "Hritik", "Aman", "Amit", "Singh", "Bisht"];
    friends.where((name) => name.startsWith("A")).forEach((name) => print(name));
  }
