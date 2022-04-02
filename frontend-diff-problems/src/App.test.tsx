import React from "react";
import { render, screen } from "@testing-library/react";
import App from "./App";

test("basic again", () => {
  expect(2).toBe(2);
});

test("renders learn react links", () => {
  render(<App />);
  const linkElement = screen.getByText(/Hello/i);
  expect(linkElement).toBeInTheDocument();
});
