import "./App.css";
import { useRef, useEffect, useState } from "react";
import MessageList from "./components/MessageList";
import MessageForm from "./components/MessageForm";

function App() {
  const ws = useRef(null);
  const [enteredMessage, setEnteredMessage] = useState("");
  const [messageList, setMessageList] = useState([]);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/ws");
  }, []);

  useEffect(() => {
    ws.current.onmessage = function (e) {
      setMessageList((prevList) => [...prevList, e.data]);
    };
  }, []);

  const submitMessageHandler = (e) => {
    e.preventDefault();
    ws.current.send(enteredMessage);
    setEnteredMessage("");
  };

  const setEnteredMessageHandler = (e) => {
    const message = e.target.value;
    setEnteredMessage(message);
  };

  return (
    <div className="chat-container">
      <div className="messages-container">
        <MessageList messageList={messageList} />
      </div>
      <MessageForm
        enteredMessage={enteredMessage}
        onSubmitMessage={submitMessageHandler}
        onChangeEnteredMessage={setEnteredMessageHandler}
      />
    </div>
  );
}

export default App;
