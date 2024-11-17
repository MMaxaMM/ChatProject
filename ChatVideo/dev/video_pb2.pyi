from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class VideoRequest(_message.Message):
    __slots__ = ("filepath",)
    FILEPATH_FIELD_NUMBER: _ClassVar[int]
    filepath: str
    def __init__(self, filepath: _Optional[str] = ...) -> None: ...

class VideoResponse(_message.Message):
    __slots__ = ("filepath",)
    FILEPATH_FIELD_NUMBER: _ClassVar[int]
    filepath: str
    def __init__(self, filepath: _Optional[str] = ...) -> None: ...
