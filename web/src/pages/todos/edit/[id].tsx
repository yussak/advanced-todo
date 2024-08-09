import { api } from "@/lib/api-client";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { SubmitHandler, useForm } from "react-hook-form";

const EditTodo = () => {
  const router = useRouter();
  const id = router.query.id;

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<Inputs>({
    mode: "onChange",
  });
  const onSubmit: SubmitHandler<Inputs> = async (data) =>
    await updateTodo(data);

  useEffect(() => {
    if (router.isReady) {
      fetchTodoData();
    }
  }, [router.isReady]);

  const fetchTodoData = async () => {
    try {
      const { data } = await api.get(`/todos/${id}`);
      setValue("title", data.title);
      setValue("body", data.body);
    } catch (error) {
      console.error(error);
    }
  };

  const updateTodo = async (data: Inputs) => {
    try {
      await api.put(`/todos/edit/${id}`, data);
      router.push(`/todos/${id}`);
    } catch (error) {
      console.error(error);
    }
  };

  type Inputs = {
    title: string;
    body: string;
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <label htmlFor="title">title</label>
        <input
          {...register("title", {
            required: "this field is required.",
          })}
        />
        {errors.title && <span>{errors.title.message}</span>}
        <label htmlFor="body">body</label>
        <input
          {...register("body", {
            required: "this field is required.",
          })}
        />
        {errors.body && <span>{errors.body.message}</span>}
        <button type="submit">update</button>
      </form>
    </>
  );
};

export default EditTodo;
