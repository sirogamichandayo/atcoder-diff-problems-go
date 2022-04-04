import React, { useState } from "react";
import "./App.css";
import { request } from "graphql-request";

const url = "https://localhost:8080/graphql";
const query = `
  query helloWorld {
    helloWorld
  }
`;

function App() {
  const [hello, setHello] = useState("");

  const graphQL = async () => {
    const result = await request(url, query);
    setHello(result.helloWorld);
  };

  const a = 10;
  return (
    <div className="App">
      {a}
      <h1>Hello</h1>
      <button onClick={graphQL}>Click</button>
      <div>{hello}</div>
    </div>
  );
}

export default App;
