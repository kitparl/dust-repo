
// 9. Write a program in Dart to remove all whitespaces from String.

void main() {
  String str = "My Name Is Pranshu Singh Bisht";

  String bag = "";
  for (int i = 0; i < str.length; i++) {
    if (str[i] != " ") {
      bag += str[i];
    }
  }
  
  print("Result: $bag");
}
