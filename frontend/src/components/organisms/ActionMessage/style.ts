import styled from "@emotion/styled";

export const ActionMessageWrapper = styled.div`
  background: rgba(59, 125, 233, 0.9);
  font-size: 13px;
  line-height: 50px;
  padding-left: 15px;
  position: relative;
  text-align: center;
  &.color-success {
    background: rgba(59, 125, 233, 0.9);
  }
  &.color-failure {
    background: rgb(245, 117, 117, 0.9);
  }
`;
