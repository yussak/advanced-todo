import { it, expect, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import TodoList from "../TodoList";
import axios from "axios";
import "@testing-library/jest-dom";

vi.mock("axios");

describe("TodoForm", () => {
  it("should display correct number of todos", async () => {
    // apiとのやり取りは親側でやっているので、TodoListコンポーネントのテストでは直書きでtodosでOK
    const todos = [
      { id: "abc", title: "title1", body: "body1" },
      { id: "def", title: "title2", body: "body2" },
    ];
    axios.get.mockResolvedValue({ data: todos });
    render(<TodoList todos={todos} />);

    const todoItems = screen.getAllByRole("listitem");
    expect(todoItems).toHaveLength(todos.length);
  });

  it("should display text when todos are empty", async () => {
    const todos = [];
    const onDelete = () => {};
    render(<TodoList todos={todos} onDelete={onDelete} />);

    expect(screen.getByText("not found")).toBeInTheDocument();
  });
});
