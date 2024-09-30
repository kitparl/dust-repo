
// 11. Suppose, you often go to restaurant with friends and you have to split amount of bill. Write a program to calculate split amount of bill. Formula= (total bill amount) / number of people

import 'dart:io';
void main(){
    print("Enter Total Bill Amount");
    int totalBill = int.parse(stdin.readLineSync()!); 

    print("Enter Number of People");
    int numberOfPeople = int.parse(stdin.readLineSync()!);

    double result = totalBill / numberOfPeople;

    print("Bill Per Person: $result");

  }
