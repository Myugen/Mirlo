import React from 'react'
import styled from 'styled-components'
import { Formik, Form, Field } from 'formik'
import * as Yup from 'yup'

import { useTranslation } from 'react-i18next'
import 'src/translations/i18n'

import FormField from './formField'

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

const SignUpForm = () => {
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
    password: Yup.string()
      .required(t('validations.required'))
      .matches(
        /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{16,128}$/,
        t('validations.passwordRequirements')
      ),
  })

  return (
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
      <Form
        style={{
          height: '100%',
          width: '100%',
        }}
      >
        <div class="w-full h-3/4 flex flex-col items-center justify-start">
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
          <Field
            name="password"
            component={FormField}
            placeholder={t('signUp.form.fields.password.example')}
            title={t('signUp.form.fields.password.title')}
          />
        </div>
        <div class="w-full h-1/4 flex items-center justify-center ">
          <Button>{t('signUp.next')}</Button>
        </div>
      </Form>
    </Formik>
  )
}

export default SignUpForm
