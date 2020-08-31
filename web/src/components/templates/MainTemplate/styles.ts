import styled from "@emotion/styled";

export const Wrapper = styled.div`
  position: relative;
  min-width: 1000px;
`;

export const Header = styled.div`
  position: fixed;
  top: 0px;
  right: 0;
  left: 0;
  z-index: 99;
`;

export const Body = styled.div`
  padding-left: 180px;
  padding-top: 40px;
  position: relative;
  z-index: 1;
  height: 100vh;
`;

export const Content = styled.div`
  padding-left: 16px;
  padding-top: 15px;
  width: 900px;
`;
