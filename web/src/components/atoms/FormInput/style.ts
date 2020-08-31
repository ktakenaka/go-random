import styled from "@emotion/styled";

export const FormInputWrapper = styled.div`
  display: contents;
  input {
    border: 1px solid #d4d8dd;
    box-sizing: border-box;
    border-radius: 4px;
    height: 28px;
    width: 100%;
    padding: 7px 8px;
    font-size: 13px;
    &.error {
      border-color: #dc3545;
      background-color: #ffeeeb;
    }
    &:focus {
      color: #495057;
      background-color: #fff;
      border-color: #80bdff;
      outline: 0;
      box-shadow: 0 0 0 0.1rem rgba(0, 123, 255, 0.25);
    }
  }
`;

export const ErrorMessage = styled.div`
  color: #dc3545;
`;
