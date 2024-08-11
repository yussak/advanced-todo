import Link from "next/link";

const TodoList = ({ todos, onDelete }) => {
  return (
    <div>
      {todos.length > 0 ? (
        todos.map((todo) => (
          <p key={todo.id} role="listitem">
            {todo.title}, {todo.body}
            <button onClick={() => onDelete(todo.id)}>delete</button>
            <Link href={`/todos/${todo.id}`}>detail</Link>
          </p>
        ))
      ) : (
        <p>not found</p>
      )}
    </div>
  );
};

export default TodoList;
