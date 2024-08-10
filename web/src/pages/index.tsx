import Head from "next/head";
import { useEffect, useState } from "react";
import Link from "next/link";
import { SubmitHandler, useForm } from "react-hook-form";
import { api } from "@/lib/api-client";

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

  const addTodo = async (data: Inputs) => {
    const { title, body } = data;
    try {
      await api.post("/todo", { title: "", body });
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
  const handleDelete = async (id: string) => {
    try {
      await api.delete(`/todo/${id}`);
      setFlashMessage("Todo deleted");
      setTimeout(() => setFlashMessage(null), 3000);
    } catch (error) {
      console.error(error);
    }
    await fetchTodos();
  };

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = async (data) => await addTodo(data);

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
          <div
            className={`${
              isError
                ? "bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative"
                : "bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative"
            }`}
            role="alert"
          >
            <strong className="font-bold">{flashMessage}</strong>
          </div>
        )}

        <h1 className="text-3xl font-bold underline">TodoList</h1>
        <form onSubmit={handleSubmit(onSubmit)} className="max-w-sm mx-auto">
          <div className="mb-5">
            <label
              htmlFor="title"
              className="block mb-2 text-sm font-medium text-gray-900"
            >
              title
            </label>
            <input
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              {...register("title", { required: "this field is required." })}
            />
            {errors.title && (
              <span className="text-red-600">{errors.title.message}</span>
            )}

            <label
              htmlFor="body"
              className="block mb-2 text-sm font-medium text-gray-900"
            >
              body
            </label>
            <input
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              {...register("body", { required: "this field is required." })}
            />
            {errors.body && (
              <span className="text-red-600">{errors.body.message}</span>
            )}
          </div>
          <button
            type="submit"
            className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Submit
          </button>
        </form>
        {todos.map((todo) => (
          <p key={todo.id}>
            {todo.title}, {todo.body}
            <button onClick={() => handleDelete(todo.id)}>delete</button>
            <Link href={`/todos/${todo.id}`}>detail</Link>
          </p>
        ))}
      </main>
    </>
  );
}
