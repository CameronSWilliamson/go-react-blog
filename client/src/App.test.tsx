import React from "react";
import { render, screen } from "@testing-library/react";
import App from "./App";
import Header from "./components/Header";

test("renders learn react link", () => {
    render(<Header />);
    const linkElement = screen.getByText(/React-Bootstrap/i);
    expect(linkElement).toBeInTheDocument();
});
