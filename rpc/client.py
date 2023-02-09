import grpc
import pythonProto.RPCtest_pb2 as myproto
import pythonProto.RPCtest_pb2_grpc  as rpctest

def run():
    with grpc.insecure_channel('localhost:'+port) as channel:
        stub = rpctest.SumStub(channel)
        response = stub.SayHello(myproto.SumNumsREQ(nums=1))
        print(response)
        response = stub.SayHello(myproto.SumNumsREQ(nums=2))
        print(response)
        


if __name__ == '__main__':
    print("Start")
    port = "18888"
    run()
    
