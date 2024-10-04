
// 12. Write a function in Dart called calculateArea that calculates the area of a rectangle. It should take length and width as arguments, with a default value of 4 for both. Formula: length * width.


void main(){
    int area = areaOfRectangle(2,4);
    print(area);
  }

areaOfRectangle([int length = 4, int width = 4]){
    return length * width;
  }
