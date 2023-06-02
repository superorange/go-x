class Person {
  Name? name;

  @override
  String toString() {
    return 'Person{name: ${name?.toString()}';
  }
}

class Name {
  String? name;

  @override
  String toString() {
    return 'Name{name: $name}';
  }

}

void main() {
  var name = Name()
    ..name = "张三";
  var person = Person()
    ..name = name;
  name.name = "李四";
  print("person:$person");
  print("name:$name");
}
