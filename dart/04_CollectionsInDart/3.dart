
//- Create a program thats reads list of expenses amount using user input and print total.

import 'dart:io';
void main(List<String> args) {
List<int> list = [];
print("Enter Expanse 1");
int num1 = int.parse(stdin.readLineSync()!);
list.add(num1);
print("Enter Expanse 2");
int num2 = int.parse(stdin.readLineSync()!);
list.add(num2);
print("Enter Expanse 3");
int num3 = int.parse(stdin.readLineSync()!);
list.add(num3);

int total = 0;
for(int i=0; i<list.length; i++){
  total += list[i];
}

print("Total is $total");
}
