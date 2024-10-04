
// 9. Write a function in Dart called maxNumber that takes three numbers as arguments and returns the largest number.

void main(){
    int max = maxNumber(4,3,2);
    print(max);
  }

  int maxNumber(int a, int b, int c){
    return a>b?a:b>c?b:c;
  }
