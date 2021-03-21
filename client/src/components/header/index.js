import React from 'react'
import styled from 'styled-components'
import { useTranslation } from 'react-i18next'
import 'src/translations/i18n'

const Title = styled.div`
  flex: 3;
  display: flex;
  justify-content: flex-start;
`

const Logout = styled.div`
  flex: 1;
  display: flex;
  justify-content: flex-end;
`

const TitleButton = styled.button`
  font-size: 1.5rem;
  text-align: center;
  color: white;
  font-family: 'Source Code Pro', monospace;
  font-weight: 600;

  &:focus {
    outline: none;
  }
`
const LogoutButton = styled.button`
  font-size: 1rem;
  text-align: center;
  color: white;
  font-family: 'Source Code Pro', monospace;

  &:focus {
    outline: none;
  }
`

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  background-color: rgba(31, 41, 55, 1);
  padding: 10px 20px 0px 20px;
  overflow: hidden;
`

const Header = () => {
  const { t } = useTranslation()

  return (
    <Wrapper>
      <Title>
        <TitleButton>{t('common.title')}</TitleButton>
      </Title>
      <Logout>
        <LogoutButton>{t('header.logOut')}</LogoutButton>
      </Logout>
    </Wrapper>
  )
}

export default Header
