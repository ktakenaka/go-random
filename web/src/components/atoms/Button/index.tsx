import React from "react";
import { ButtonWrapper } from "./style";

type Props = {
  size: "shorter" | "short" | "medium" | "big";
  color: "blue" | "grey";
  disabled: boolean;
  type: "submit" | "button";
  value: string;
};

const Button = ({ type, size, color, disabled, value }: Props) => {
  return (
    <ButtonWrapper
      className={`color-${color} size-${size}`}
      type={type}
      disabled={disabled}
    >
      {value}
    </ButtonWrapper>
  );
};

export default Button;
