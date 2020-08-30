import React from 'react';
import { Button as ButtonAnt } from 'antd';
import { ButtonWrapper } from './style'

type Props = {
  size: 'shorter' | 'short' | 'medium' | 'big',
  color: 'blue' | 'grey',
  children: React.ReactNode,
  disabled: boolean,
  type: 'button' | 'submit',
  onClick:((...params: any[]) => any)
};

const Button = ({ type, size, color, children, disabled, onClick }:Props) => {
  return (
    <ButtonWrapper
      className={`color-${color} size=${size}`}
      onClick={onClick}
      type={type}
      disabled={disabled}
    >
      <ButtonAnt>
        {children}
      </ButtonAnt>
    </ButtonWrapper>
  )
};

export default Button
