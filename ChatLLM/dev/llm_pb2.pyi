from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Message(_message.Message):
    __slots__ = ("role", "content")
    ROLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    role: str
    content: str
    def __init__(self, role: _Optional[str] = ..., content: _Optional[str] = ...) -> None: ...

class LLMRequest(_message.Message):
    __slots__ = ("messages", "max_tokens")
    MESSAGES_FIELD_NUMBER: _ClassVar[int]
    MAX_TOKENS_FIELD_NUMBER: _ClassVar[int]
    messages: _containers.RepeatedCompositeFieldContainer[Message]
    max_tokens: int
    def __init__(self, messages: _Optional[_Iterable[_Union[Message, _Mapping]]] = ..., max_tokens: _Optional[int] = ...) -> None: ...

class LLMResponse(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: Message
    def __init__(self, message: _Optional[_Union[Message, _Mapping]] = ...) -> None: ...
