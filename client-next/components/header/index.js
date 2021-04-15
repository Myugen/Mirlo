import React from 'react'
import styled from 'styled-components'
import { useTranslation } from 'next-i18next'
import { useRouter } from 'next/dist/client/router'

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

const TitleHeader = styled.a`
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

const Wrapper = styled.header`
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 10px 20px 0px 20px;
  overflow: hidden;
`

const Header = () => {
  const { t } = useTranslation()
  const router = useRouter()

  const handleTitleClick = (e) => {
    e.preventDefault()
    router.push('/')
  }

  return (
    <Wrapper>
      <Title>
        <TitleHeader href="/" onClick={handleTitleClick}>
          {t('title')}
        </TitleHeader>
      </Title>
      <Logout>
        <LogoutButton>{t('header.logOut')}</LogoutButton>
      </Logout>
    </Wrapper>
  )
}

export default Header
