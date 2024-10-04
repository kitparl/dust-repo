
// 2. Write a program in Dart to print even numbers between intervals using function.


void main(){
    evenNumber(start: 1, end: 10);
  }

  void evenNumber({int? start, int? end}){

    for(int i=start!; i<end!; i++){
      if(i%2==0){
        print(i);
      }
    }
  }

