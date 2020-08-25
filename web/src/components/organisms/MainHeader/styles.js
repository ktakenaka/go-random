import styled from '@emotion/styled';

export const Wrapper = styled.div`
  display: flex;
  flex-flow: row nowrap;
  padding: 8px 0px 12px 20px;
  color: #3A3A3A;
  justify-content: space-between;
  height: 40px;
  border-bottom: 1px solid #D4D8DD;
  background-color: #FFFFFF;
  svg {
    color: #7C8291;
  }
`;

export const Devide = styled.div`
  height: 18px;
  border-left: 1px solid #d8d8d8;
  margin: 0 11px 0 18px;
`;

export const LeftContent = styled.div`
  display: flex;
  box-sizing: content-box;
  padding-top: 2px;
`;

export const RightContent = styled.div`
  display: flex;
`;

export const Item = styled.div`
  margin-right: 32px;
`;
