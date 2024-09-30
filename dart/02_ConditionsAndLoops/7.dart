// 7. Write a dart program to generate multiplication tables of 1-9.


void main(){
  for(int i=1; i<9; i++){
    int num = i;
    print("********$num table********");
    for(int j=1; j<10; j++){
      print("$num*$j = ${num*j}");
    }
    print("********************");
  }
  }
