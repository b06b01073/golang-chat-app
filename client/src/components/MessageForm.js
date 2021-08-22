import React from "react";

const MessageForm = (props) => {
  const enteredMessage = props.enteredMessage;
  const submitMessageHandler = props.onSubmitMessage;
  const setEnteredMessageHandler = props.onChangeEnteredMessage;

  return (
    <div>
      <form onSubmit={submitMessageHandler}>
        <input
          type="text"
          placeholder="Write your message here..."
          value={enteredMessage}
          onChange={setEnteredMessageHandler}
        ></input>
        <button type="submit">Send</button>
      </form>
    </div>
  );
};

export default MessageForm;
