
// - Create an empty list of type string called days. Use the add method to add names of 7 days and print all days.

void main(){
    List<String> days = [];
    days.add("Mon");
    days.add("Tue");
    days.add("Wed");
    days.add("Thu");
    days.add("Fri");
    days.add("Sat");
    days.add("Sun");
    for(String day in days){
      print(day);
    }
  }
