import React from "react";

export default function App({ message }) {
  return (
    <div className="container">
      <ul className="nav">
        <li className="nav-item">
          <a className="nav-link active" href="/">
            Home
          </a>
        </li>
        <li className="nav-item">
          <a className="nav-link" href="/racers">
            Racers
          </a>
        </li>
      </ul>
      <h1>{message}</h1>
    </div>
  );
}
