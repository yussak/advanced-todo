import { SubmitHandler, useForm } from "react-hook-form";

type Inputs = {
  title: string;
  body: string;
};

interface TodoFormProps {
  onSubmit: (data: Inputs, reset: () => void) => void;
}

const TodoForm: React.FC<TodoFormProps> = ({ onSubmit }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<Inputs>();

  const handleFormSubmit: SubmitHandler<Inputs> = async (data) => {
    await onSubmit(data, reset);
  };

  return (
    <>
      <form
        onSubmit={handleSubmit(handleFormSubmit)}
        className="max-w-sm mx-auto"
      >
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
        {/* TODO:エラーある時押せなくする */}
        <button
          type="submit"
          className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        >
          Submit
        </button>
      </form>
    </>
  );
};

export default TodoForm;
