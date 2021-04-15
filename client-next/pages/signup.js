import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import styled from 'styled-components'
import { useTranslation } from 'next-i18next'
import Header from 'components/header'
import { SignUpForm } from 'components/signup'

export const getStaticProps = async ({ locale }) => ({
  props: {
    ...(await serverSideTranslations(locale, [
      'common',
      'header',
      'validations',
      'signup',
    ])),
  },
})

const Title = styled.h1`
  font-family: 'Source Code Pro', monospace;
  font-size: 35px;
  font-weight: 600;
  color: white;
`

const SignUp = () => {
  const { t } = useTranslation('signup')

  return (
    <div className="w-screen h-screen bg-gradient-to-br from-gray-800 to-gray-700">
      <div className="w-screen h-auto absolute">
        <Header />
      </div>
      <div className="w-full h-full flex justify-center items-center">
        <div className="w-1/5 h-3/4 bg-gray-900 rounded-xl shadow-2xl">
          <div className="w-full h-1/4 flex justify-center items-center">
            <Title>{t('title')}</Title>
          </div>
          <div className="w-full h-3/4 flex flex-col justify-start items-center">
            <SignUpForm />
          </div>
        </div>
      </div>
    </div>
  )
}

export default SignUp
