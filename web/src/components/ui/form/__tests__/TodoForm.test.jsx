import { it, expect, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import TodoList from "../../../todos/TodoList";
import axios from "axios";

vi.mock("axios");

it("個数が正しい", async () => {
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
