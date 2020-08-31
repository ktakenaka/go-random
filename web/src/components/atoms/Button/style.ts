import styled from '@emotion/styled';

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
    background: linear-gradient(180deg, #3B7DE9 0%, #0054AC 100%) !important;
    border: 1px solid #3B7DE9 !important;
    color: #FFFFFF !important;
    &:hover {
      background: linear-gradient(90deg, #3B7DE9 0%, #0054AC 100%) !important;
    }
  }
  &.color-grey {
    background: linear-gradient(180deg, #FFFFFF 0%, #EFF1F4 100%);
    border: 1px solid #D4D8DD;
    color: #3B7DE9;
    &:hover {
      background: linear-gradient(90deg, #FFFFFF 0%, #EFF1F4 100%);
    }
  }
`;
