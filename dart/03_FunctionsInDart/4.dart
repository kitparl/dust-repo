
// 4. Write a program in Dart that generates random password.

import "dart:math";
void main(){
  print(password());
}

int password(){
int randomnum = 10000 + Random().nextInt((20000 + 1) - 10000);

return randomnum;
}
