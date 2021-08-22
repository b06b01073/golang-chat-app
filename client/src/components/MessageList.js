import React from "react";
import Message from "./Message";

const MessageList = (props) => {
  const messageList = props.messageList.map((message, index) => (
    <Message message={message} key={index} />
  ));

  return <>{messageList}</>;
};

export default MessageList;
