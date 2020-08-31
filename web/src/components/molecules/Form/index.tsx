import React from "react";

type Props = {
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
  children: React.ReactNode;
};

const Form = ({ children, onSubmit }: Props) => {
  return <form onSubmit={onSubmit}>{children}</form>;
};

export default Form;
