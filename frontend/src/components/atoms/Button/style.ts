import styled from "@emotion/styled";

export const ButtonWrapper = styled.button`
  font-size: 14px;
  font-weight: 500;
  line-height: 20px;
  box-sizing: border-box;
  border-radius: 4px;
  cursor: pointer;
  padding: 4px 0px;
  &:disabled {
    opacity: 0.5;
  }
  &.size-shorter {
    min-width: 90px;
    height: 32px;
  }
  &.size-short {
    min-width: 100px;
    height: 32px;
  }
  &.size-medium {
    min-width: 130px;
    height: 32px;
  }
  &.size-big {
    min-width: 240px;
    height: 44px;
  }
  &.color-blue {
    background: linear-gradient(180deg, #3b7de9 0%, #0054ac 100%) !important;
    border: 1px solid #3b7de9 !important;
    color: #ffffff !important;
    &:hover {
      background: linear-gradient(90deg, #3b7de9 0%, #0054ac 100%) !important;
    }
  }
  &.color-grey {
    background: linear-gradient(180deg, #ffffff 0%, #eff1f4 100%);
    border: 1px solid #d4d8dd;
    color: #3b7de9;
    &:hover {
      background: linear-gradient(90deg, #ffffff 0%, #eff1f4 100%);
    }
  }
`;
