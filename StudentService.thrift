namespace go go_src.go_stu
namespace py py_src.py_stu

struct Student {
    1: string name;
    2: i32 age;
    3: set<i32> scores;
}

service StudentService {
    map<string, list<Student>> groupByName(1:list<Student> students); 

    i32 length(1:string str);

    set<i32> distinct(1:list<i32> values);
}