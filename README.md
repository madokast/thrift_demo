# thrift_demo

## code

StudentService.thrift

```
namespace go go_stu
namespace py py_stu

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
```

## gen

`thrift-0.17.0.exe --gen py .\StudentService.thrift`
`thrift-0.17.0.exe --gen go .\StudentService.thrift`

## python

`pip install thrift -i https://pypi.tuna.tsinghua.edu.cn/simple`

