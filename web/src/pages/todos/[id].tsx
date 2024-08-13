import { api } from "@/lib/api-client";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

// TODO:共通化
type Todo = {
  id: string;
  title: string;
  body: string;
};

const TodoDetail = () => {
  const router = useRouter();
  const id = router.query.id;
  const [todo, setTodo] = useState<Todo | null>(null);

  useEffect(() => {
    if (id) {
      getGoalDetails(id as string);
    }
  }, [id]);

  const getGoalDetails = async (id: string) => {
    try {
      const res = await api.get(`/todos/${id}`);
      setTodo(res.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      {todo ? (
        <>
          <p>title: {todo.title}</p>
          <p>body: {todo.body}</p>

          <Link href={`/todos/edit/${todo.id}`}>edit</Link>
        </>
      ) : (
        // TODO:一瞬not foundがちらつくので要修正→共通のローディング画面を作るので対応したい
        <p>todo not found</p>
      )}
    </>
  );
};

export default TodoDetail;
