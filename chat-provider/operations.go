package main

type Operation string

const (
	//JOIN_CHANNEL_OPERATION  Operation = "Wcchat.JoingChannel"
	//LEAVE_CHANNEL_OPERATION Operation = "Wcchat.LeaveChannel"
	//SEND_MSG_OPERATION      Operation = "Wcchat.SendMsg"
	//RECEIVE_MSG_OPERATION   Operation = "Wcchat.ReceiveMsg"

	UPDATE_USER_LIST Operation = "Chat.UpdateUserList"
	UPDATE_CHAT      Operation = "Chat.UpdateChat"

	SEND_CHAT_MSG    Operation = "ChatSubscriber.HandleChatMessage"
	SEND_USER_UPDATE Operation = "ChatSubscriber.HandleUserMessage"
)

func (o Operation) String() string {
	return string(o)
}
