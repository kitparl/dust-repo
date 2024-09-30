
// 8. Write a dart program to create a simple calculator that performs addition, subtraction, multiplication, and division.

import 'dart:io';
void main(){
  print("Enter first number: ");
  int num1 = int.parse(stdin.readLineSync()!);
  print("Enter second number: ");
  int num2 = int.parse(stdin.readLineSync()!);
  print("Select Operation : 1 -> additon; 2 -> subtraction; 3 -> Multiplication; 4 -> Division");
  int operation = int.parse(stdin.readLineSync()!);
  switch(operation){
    case 1:
      print("$num1 + $num2 = ${num1+num2}");
      break; 
    case 2:
      print("$num1 - $num2 = ${num1-num2}");
      break; 
    case 3:
      print("$num1 * $num2 = ${num1*num2}");
      break; 
    case 4:
      print("$num1 / $num2 = ${num1/num2}");
      break; 
  }
}
