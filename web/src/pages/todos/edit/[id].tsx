import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const EditTodo = () => {
  const router = useRouter();
  const id = router.query.id;
  const [todoData, setTodoData] = useState({
    title: "",
    body: "",
  });

  useEffect(() => {
    if (router.isReady) {
      fetchTodoData();
    }
  }, [router.isReady]);

  const fetchTodoData = async () => {
    try {
      const { data } = await axios.get(`http://localhost:8080/todos/${id}`);
      setTodoData({ title: data.title, body: data.body });
    } catch (error) {
      console.error(error);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setTodoData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.put(`http://localhost:8080/todos/edit/${id}`, todoData);
      router.push(`/todos/${id}`);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      <form onSubmit={handleSubmit}>
        <label htmlFor="title">title</label>
        <input
          type="text"
          name="title"
          value={todoData.title}
          onChange={handleChange}
        />

        <label htmlFor="body">body</label>
        <input
          type="text"
          name="body"
          value={todoData.body}
          onChange={handleChange}
        />

        <button type="submit">update</button>
      </form>
    </>
  );
};

export default EditTodo;
