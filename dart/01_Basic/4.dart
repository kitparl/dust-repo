void main(List<String> args) {
  int principle = 300;
  int rate = 12;
  int time = 2;

  double simpleInterest = (principle * rate * time) / 100;

  print(simpleInterest.toInt()); // toInt() is similar to truncate() i am using to convert my double to int
}
