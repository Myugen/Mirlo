import React from 'react'
import styled from 'styled-components'
import { Formik, Form, Field } from 'formik'
import * as Yup from 'yup'
import useSwr from 'swr'
import { useTranslation } from 'next-i18next'

import { useStore } from 'store'
import { ApiErrorAlert } from '../alert'
import { FormField } from '../form'

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
  const { t: signupT } = useTranslation('signup')
  const { t: validationsT } = useTranslation('validations')

  const { dispatch } = useStore

  const SignupSchema = Yup.object().shape({
    username: Yup.string()
      .min(2, validationsT('tooShort'))
      .max(20, validationsT('tooLong'))
      .required(validationsT('required')),
    email: Yup.string()
      .min(3, validationsT('tooShort'))
      .max(320, validationsT('tooLong'))
      .required(validationsT('required')),
    password: Yup.string()
      .min(16, validationsT('tooShort'))
      .max(128, validationsT('tooLong'))
      .required(validationsT('required'))
      .matches(
        /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{16,128}$/,
        validationsT('passwordRequirements')
      ),
  })

  const submitForm = (values) => {
    const { username, password, email } = values
    axios
      .post('/users', {
        username,
        password,
        email,
      })
      .then(({ data }) => {
        dispatch({
          type: 'SET_USER',
          payload: {
            user: data,
          },
        })
      })
      .catch((err) => {
        ApiErrorAlert('Error!', err.response.data.message)
      })
  }

  return (
    <Formik
      initialValues={{
        username: '',
        email: '',
        password: '',
      }}
      validationSchema={SignupSchema}
      onSubmit={(values) => {
        submitForm(values)
      }}
    >
      <Form
        style={{
          height: '100%',
          width: '100%',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <div className="w-3/4 h-3/4 flex flex-col items-center justify-start">
          <Field
            name="username"
            component={FormField}
            placeholder={signupT('form.fields.username.example')}
            title={signupT('form.fields.username.title')}
          />
          <Field
            name="email"
            component={FormField}
            placeholder={signupT('form.fields.email.example')}
            title={signupT('form.fields.email.title')}
          />
          <Field
            name="password"
            component={FormField}
            placeholder={signupT('form.fields.password.example')}
            title={signupT('form.fields.password.title')}
          />
        </div>
        <div className="w-full h-1/4 flex items-center justify-center">
          <Button>{signupT('next')}</Button>
        </div>
      </Form>
    </Formik>
  )
}

export default SignUpForm
