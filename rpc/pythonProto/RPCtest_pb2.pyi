from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class SumNumsACK(_message.Message):
    __slots__ = ["result"]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: int
    def __init__(self, result: _Optional[int] = ...) -> None: ...

class SumNumsREQ(_message.Message):
    __slots__ = ["nums"]
    NUMS_FIELD_NUMBER: _ClassVar[int]
    nums: int
    def __init__(self, nums: _Optional[int] = ...) -> None: ...
