
// 10. Write a function in Dart called isEven that takes a number as an argument and returns True if the number is even, and False otherwise.

void main(){
    int num = 5;
    bool result = isEven(num);
    print(result);
  }


  bool isEven(int num){
    return num%2==0;
  }
