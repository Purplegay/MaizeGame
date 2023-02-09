'''
@Author: EndLess
@Date: 2019-08-09 14:07:32
@LastEditors: Please set LastEditors
@LastEditTime: 2020-03-17 18:28:04
@Description: file content
'''
import os

all_file=""

for dir,j,files in os.walk("./"):
    for file in files:
        ext=os.path.splitext(file)[-1]
        if ext==".proto":
            all_file+=file+" "

# cmd ="protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "+all_file
# print(cmd)
# os.system(cmd)

cmd ="python -m grpc_tools.protoc -I. --python_out=../rpc/pythonProto --pyi_out=../rpc/pythonProto --grpc_python_out=../rpc/pythonProto "+all_file
print(cmd)
os.system(cmd)