const FlashMessage = ({
  isError,
  flashMessage,
}: {
  isError: boolean;
  flashMessage: string;
}) => {
  return (
    <>
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
    </>
  );
};

export default FlashMessage;
