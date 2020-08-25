import React from 'react';
import { Wrapper, LeftContent, RightContent, Devide, Item } from './styles';

const MainHeader = () => {
  return (
    <Wrapper>
      <LeftContent>
        <Devide />
        <Item>Left</Item>
      </LeftContent>
      <RightContent>
        <Item>Right</Item>
      </RightContent>
    </Wrapper>
  )
}

export default MainHeader;
