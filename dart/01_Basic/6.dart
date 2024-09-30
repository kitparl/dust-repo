// 6. Write a program to print full name of a from first name and last name using user input.
import 'dart:io'; 
void main(List<String> args) {
  print("Enter Your Full Name");
  String? name = stdin.readLineSync();

  var list = name?.split(" ");


  String firstName = list![0];
  print(firstName);

  String lastName = list[list.length - 1];
  print(lastName);
}
