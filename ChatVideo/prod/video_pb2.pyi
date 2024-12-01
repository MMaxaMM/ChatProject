from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class VideoRequest(_message.Message):
    __slots__ = ("URI",)
    URI_FIELD_NUMBER: _ClassVar[int]
    URI: str
    def __init__(self, URI: _Optional[str] = ...) -> None: ...

class VideoResponse(_message.Message):
    __slots__ = ("objectName",)
    OBJECTNAME_FIELD_NUMBER: _ClassVar[int]
    objectName: str
    def __init__(self, objectName: _Optional[str] = ...) -> None: ...
