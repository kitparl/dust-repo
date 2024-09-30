// 5. Write a program to print a square of a number using user input.
import 'dart:io';
void main(List<String> args) {
 print("Enter Number for sqaure"); 

int num = int.parse(stdin.readLineSync()!);
// or alternatively

/*
String? input = stdin.readLineSync();
if (input != null) {
  int? number = int.parse(input);
}
*/

print(num * num);
}
