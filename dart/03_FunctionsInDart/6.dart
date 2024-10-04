
// 6. Write a program in Dart to reverse a String using function.

void main(){
    var rev = reverse();
    print(rev);
  }

  String reverse(){
    String name = "Pranshu";
    String bag = "";
    for(int i=name.length-1; i>=0; i--){
      bag += name[i];
    }
    return bag;
  }
