metadata package = [ { namespace: "org.jordanrash.chat" } ]

namespace org.jordanrash.chat

use org.wasmcloud.model#wasmbus

@wasmbus(
    contractId: "jordanrash:wcchat",
    providerReceive: true )
service Chat {
  version: "0.1",
  operations: [ UpdateUserList, UpdateChat ]
}

@wasmbus(
    contractId: "jordanrash:wcchat",
    actorReceive: true )
service ChatSubscriber {
  version: "0.1",
  operations: [ HandleChatMessage, HandleUserMessage ]
}

operation UpdateUserList {
    input:  UserChannelMsg,
}

operation UpdateChat {
    input:  Msg,
}

operation HandleChatMessage {
    input: Msg,
}

operation HandleUserMessage {
    input: Msg,
}

structure User {
    name: String,
    color: String
}

structure Msg {
    time:  String,
    value: String,
    owner: User,
    id: String
}

structure UserChannelMsg {
    action: String,
    user: User 
}
