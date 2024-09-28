
// 10. Write a Dart program to convert String to int.

void main(List<String> args) {
  String val = "12";
  print(val.runtimeType);
  int res = int.parse(val);
  print(res.runtimeType);
  print("res: $res");
}
