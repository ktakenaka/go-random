import styled from "@emotion/styled";

export const ActionMessageWrapper = styled.div`
  font-size: 13px;
  line-height: 50px;
  position: relative;
  text-align: center;
  &.color-success {
    background: blue;
  }
  &.color-failure {
    background: pink;
  }
`;
