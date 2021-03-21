import React from 'react'
import styled from 'styled-components'
import PropTypes from 'prop-types'
import 'src/translations/i18n'

const Wrapper = styled.div`
  padding: 0.5em;
`

const Label = styled.h1`
  font-family: 'Source Code Pro', monospace;
  font-size: 1em;
  color: white;
`

const ErrorMessage = styled.h1`
  font-family: 'Source Code Pro', monospace;
  font-size: 0.9em;
  font-weight: 600;
  color: red;
`

const Input = styled.input`
  padding: 0.5em;
  margin-top: 0.5em;
  font-family: 'Source Code Pro', monospace;
  background: white;
  border: none;
  border-radius: 3px;

  &:focus {
    outline: none;
  }
`
const FormField = ({ field, form: { errors }, title, ...props }) => {
  return (
    <Wrapper>
      <Label>{title}:</Label>
      <Input type="text" {...props} {...field} />
      {errors[field.name] ? (
        <ErrorMessage>{errors[field.name]}</ErrorMessage>
      ) : null}
    </Wrapper>
  )
}

FormField.propTypes = {
  field: PropTypes.object,
  title: PropTypes.string,
  form: PropTypes.object,
}

export default FormField
