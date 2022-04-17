import React from 'react';
import "./App.css";

function App() {
  const hello = 'hello';

  const a = 10;
  return (
    <div className="App">
      {a}
      <h1>Hello</h1>
      <div>{hello}</div>
    </div>
  );
}

export default App;
