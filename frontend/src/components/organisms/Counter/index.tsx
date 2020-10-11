import React from "react";
import { Button, Typography } from "antd";

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
    <>
      <Typography.Title level={5}>Clicked: {value} times</Typography.Title>
      <Button type="primary" onClick={onIncrementAsync}>
        Increment after 1 second
      </Button>
      <br />
      <Button onClick={onIncrement}>Increment</Button>{" "}
      <Button onClick={onDecrement}>Decrement</Button>
    </>
  );
};

export default Counter;
