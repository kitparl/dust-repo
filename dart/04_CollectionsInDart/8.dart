
// - Create a simple to-do application that allows user to add, remove, and view their task.

import 'dart:io';
void main(){
    
    List<String> list = [];
    print("Todo application");
    while(true){
      print("1. Add | 2. Remove | 3. View | 4. Exit");
      int input = int.parse(stdin.readLineSync()!);
      if (input == 1) {
        String work = stdin.readLineSync()!;
        list.add(work);
        print("Added Successfully");
      }else if(input == 2){
        String removeItem = stdin.readLineSync()!;
        int index = list.indexOf(removeItem);
        if(index == -1){
          print("Not Found");
        }else{
          list.removeAt(index);
          print("Removed $removeItem");
        }
      }else if(input == 3){
        list.forEach((element) {print(element);});
      }else{
        break;
      }
    }}
