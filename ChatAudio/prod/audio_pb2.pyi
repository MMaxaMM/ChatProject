from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AudioRequest(_message.Message):
    __slots__ = ("URI",)
    URI_FIELD_NUMBER: _ClassVar[int]
    URI: str
    def __init__(self, URI: _Optional[str] = ...) -> None: ...

class AudioResponse(_message.Message):
    __slots__ = ("content",)
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    content: str
    def __init__(self, content: _Optional[str] = ...) -> None: ...
