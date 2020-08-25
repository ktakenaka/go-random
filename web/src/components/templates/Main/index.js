import React from 'react';
import PropTypes from 'prop-types'

import { Wrapper, Header, Body, Content } from './styles';
import MainHeader from '../../organisms/MainHeader/index'


const MainTemplate = ({ header, children }) => {
  return (
    <Wrapper>
      <Header>{ header }</Header>
      <Body>
        <Content>{children }</Content>
      </Body>
    </Wrapper>
  )
}

MainTemplate.propTypes = {
  header: PropTypes.object,
  children: PropTypes.object,
}

MainTemplate.defaultProps = {
  header: <MainHeader />,
}

export default MainTemplate
