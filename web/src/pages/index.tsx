import Head from "next/head";
import { useEffect, useState } from "react";
import { api } from "@/lib/api-client";
import TodoList from "@/components/todos/TodoList";
import TodoForm from "@/components/ui/form/TodoForm";
import FlashMessage from "@/components/ui/flash/FlashMessage";

type Inputs = {
  title: string;
  body: string;
};

export default function Home() {
  const [todos, setTodos] = useState([]);
  const [flashMessage, setFlashMessage] = useState<string | null>(null);
  const [isError, setIsError] = useState<boolean>(false);

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const res = await api.get("/todos");
      setTodos(res.data);
    } catch (error) {
      console.error(error);
    }
  };

  const handleAddTodo = async (data: Inputs, reset: () => void) => {
    const { title, body } = data;
    try {
      await api.post("/todo", { title, body });
      setFlashMessage("Todo added");
      setIsError(false);
      setTimeout(() => setFlashMessage(null), 3000);
    } catch (error) {
      // バックエンド側のエラーを受け取ってフラッシュに出す
      if (error.response && error.response.status === 400) {
        setFlashMessage(
          error.response.data.error || "An unexpected error occurred"
        );
        setTimeout(() => setFlashMessage(null), 5000);
        setIsError(true);
      } else {
        setFlashMessage("An unexpected error occurred");
        setTimeout(() => setFlashMessage(null), 5000);
        setIsError(true);
      }
      console.error(error);
    }
    await fetchTodos();
    reset();
  };

  // TODO:サーバーのエラー受け取れるようにする
  const handleDeleteTodo = async (id: string) => {
    try {
      await api.delete(`/todo/${id}`);
      setFlashMessage("Todo deleted");
      setTimeout(() => setFlashMessage(null), 3000);
    } catch (error) {
      console.error(error);
    }
    await fetchTodos();
  };

  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="bg-gray-200">
        {/* TODO:出る時ずれるので修正 */}
        {flashMessage && (
          <FlashMessage isError={isError} flashMessage={flashMessage} />
        )}

        <h1 className="text-3xl font-bold underline">TodoList</h1>
        <TodoForm onSubmit={handleAddTodo} />
        <TodoList todos={todos} onDelete={handleDeleteTodo} />
      </main>
    </>
  );
}
