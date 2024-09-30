// 7. Write a program to find quotient and remainder of two integers.

void main(){
    int num1 = 5;
    int num2 = 2;

    int remainder = num1 % num2;
    int quotient = num1 ~/ num2; // ~/ if two operands are int, ~/ gives int quotient
    double quotient$ = num1 / num2; // if two operands are int, / gives double quotient

    print("Remainder: $remainder Quotient: $quotient");

  }
