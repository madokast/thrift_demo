from itertools import groupby
from typing import Set, List, Dict
from py_stu import StudentService
from py_stu.ttypes import Student
from thrift.transport import TSocket, TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer

# 实现逻辑
class StudentServiceImpl(StudentService.Iface):
    def groupByName(self, students:List[Student]):
        print(f"[Py_Server]groupByName students = {students}. Type = {type(students)}. EType = {type(students[0])}") # [Py_Server]groupByName students = [Student(name='mdk', age=14, scores={2}), Student(name='mdk', age=14, scores={3}), Student(name='zrx', age=15, scores={78})]. Type = <class 'list'>. EType = <class 'py_stu.ttypes.Student'>
        g = groupby(students, key=lambda s:s.name)
        return {n:list(ss) for n, ss in g}

    def length(self, str:str)->int:
        print(f"[Py_Server]length str = {str}. Type = {type(str)}") # [Py_Server]length str = abc. Type = <class 'str'>
        return len(str)

    def distinct(self, values:List[int])->Set[int]: 
        print(f"[Py_Server]distinct values = {values}. Type = {type(values)}. EType = {type(values[0])}") # [Py_Server]distinct values = [1, 1, 2, 2, 3]. Type = <class 'list'>. EType = <class 'int'>
        return set(values)
# 服务端
def start_server(port:int):
    processor = StudentService.Processor(StudentServiceImpl())
    transport = TSocket.TServerSocket(port=port)
    transFactory = TTransport.TBufferedTransportFactory()
    protocolFactory = TBinaryProtocol.TBinaryProtocolFactory()
    server = TServer.TSimpleServer(processor, transport, transFactory, protocolFactory)
    print("[Py_Server]starting")
    server.serve()
# 客户端
def start_client(port:int)->StudentService.Client:
    socket = TSocket.TSocket(host="localhost", port=port)
    transport = TTransport.TBufferedTransport(socket)
    protocol = TBinaryProtocol.TBinaryProtocol(transport)
    client = StudentService.Client(protocol)
    transport.open()
    # transport.close()
    return client

if __name__ == '__main__':
    from multiprocessing import Process
    import time
    p = Process(target=start_server, args=(18081,), daemon=True)
    p.start()
    time.sleep(1)

    client = start_client(18081)
    len_r = client.length("abc")
    print(f"[Py_Client]length return = {len_r}. Type = {type(len_r)}") # [Py_Client]length return = 3. Type = <class 'int'>

    dis_r = client.distinct([1,1,2,2,3])
    print(f"[Py_Client]distinct return = {dis_r}. Type = {type(dis_r)} {type(list(dis_r)[0])}") # [Py_Client]distinct return = {1, 2, 3}. Type = <class 'set'> <class 'int'>

    gro_r = client.groupByName([Student("mdk", 14, set([2,3])), Student("mdk", 14, set([3,4])), Student("zrx", 15, set([78]))]) # [Py_Client]groupByName return = {'mdk': [Student(name='mdk', age=14, scores={2}), Student(name='mdk', age=14, scores={3})], 'zrx': [Student(name='zrx', age=15, scores={78})]}. Type = <class 'dict'>
    print(f"[Py_Client]groupByName return = {gro_r}. Type = {type(gro_r)}")