import React from 'react'
import Header from 'src/components/header'
import styled from 'styled-components'
import { Formik, Form, Field } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import 'src/translations/i18n'

import FormField from './formField'

const Title = styled.h1`
  font-family: 'Source Code Pro', monospace;
  font-size: 35px;
  font-weight: 600;
  color: white;
`

const Button = styled.button`
  padding: 0.3em 1em 0.3em 1em;
  border-radius: 0.3em;
  font-size: 1.3em;
  font-family: 'Source Code Pro', monospace;
  font-weight: 600;
  text-align: center;
  color: rgba(17, 24, 39, 1);
  background-color: rgba(249, 250, 251, 1);

  &:focus {
    outline: none;
  }
`

const SignUp = () => {
  const { t } = useTranslation()

  const SignupSchema = Yup.object().shape({
    username: Yup.string()
      .min(2, t('validations.tooShort'))
      .max(20, t('validations.tooLong'))
      .required(t('validations.required')),
    email: Yup.string()
      .min(3, t('validations.tooShort'))
      .max(320, t('validations.tooLong'))
      .required(t('validations.required')),
  })

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
          <div class="w-full h-3/4 flex flex-col justify-center items-center">
            <div class="w-3/4 h-3/4 p-4 flex flex-col justify-center items-center">
              <Formik
                initialValues={{
                  username: '',
                  email: '',
                  password: '',
                }}
                validationSchema={SignupSchema}
                onSubmit={(values) => {
                  console.log('Hi')
                  console.log(values)
                }}
              >
                <Form>
                  <Field
                    name="username"
                    component={FormField}
                    placeholder={t('signUp.form.fields.username.example')}
                    title={t('signUp.form.fields.username.title')}
                  />
                  <Field
                    name="email"
                    component={FormField}
                    placeholder={t('signUp.form.fields.email.example')}
                    title={t('signUp.form.fields.email.title')}
                  />
                </Form>
              </Formik>
            </div>
            <div class="w-3/4 h-1/4 flex justify-center items-center">
              <Button type="submit">{t('signUp.next')}</Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default SignUp
