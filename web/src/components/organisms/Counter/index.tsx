import React, { Fragment } from "react";

type Props = {
  value: number;
  onIncrement: (e: React.MouseEvent<HTMLButtonElement>) => void;
  onDecrement: (e: React.MouseEvent<HTMLButtonElement>) => void;
  onIncrementAsync: (e: React.MouseEvent<HTMLButtonElement>) => void;
};

const Counter = ({
  value,
  onIncrement,
  onDecrement,
  onIncrementAsync,
}: Props) => {
  return (
    <Fragment>
      <button onClick={onIncrementAsync}>Increment after 1 second</button>{" "}
      <button onClick={onIncrement}>Increment</button>{" "}
      <button onClick={onDecrement}>Decrement</button>
      <hr />
      <div>Clicked: {value} times</div>
    </Fragment>
  );
};

export default Counter;
