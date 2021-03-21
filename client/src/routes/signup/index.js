import React from 'react'
import Header from 'src/components/header'
import styled from 'styled-components'

import { useTranslation } from 'react-i18next'
import 'src/translations/i18n'

import SignUpForm from './signUpForm'

const Title = styled.h1`
  font-family: 'Source Code Pro', monospace;
  font-size: 35px;
  font-weight: 600;
  color: white;
`

const SignUp = () => {
  const { t } = useTranslation()

  return (
    <div class="w-screen h-screen bg-gradient-to-br from-gray-800 to-gray-700">
      <div class="w-screen h-auto absolute">
        <Header />
      </div>
      <div class="w-full h-full flex justify-center items-center">
        <div class="w-1/5 h-3/5 bg-gray-900 rounded-xl shadow-2xl">
          <div class="w-full h-1/4 flex justify-center items-center">
            <Title>{t('signUp.title')}</Title>
          </div>
          <div class="w-full h-3/4 flex flex-col justify-start items-center">
            <SignUpForm />
          </div>
        </div>
      </div>
    </div>
  )
}

export default SignUp
