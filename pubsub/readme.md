Problem Statement
https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/pub-sub-system.md

Publisher
        | - Topic -- Message
Subscriber

Broker
    Id int
    Topics map[int]*Topic
    Subscriber map[int]*Subscriber
    mu sync.RWMutex

Topic
    Id int
    Name string
    CreatedAt int64
    Subscribers map[int]bool

Subscriber
    Id int
    Name string
    MsgChan chan *Message

Message
    Id int
    Payload string
    PublishedAt int64

Publisher
    Id int
    Name string

