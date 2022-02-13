import React from "react";

export default function Header({ message }) {
  return (
    <header className="bg-white text-white">
      <div className="container text-center text-dark">
        <h1>{message}</h1>
      </div>
    </header>
  );
}
