import React, { createContext, useContext, useReducer } from 'react'
import PropTypes from 'prop-types'
import { userActions, userInitialState } from './actions/user'

const initialState = {
  ...userInitialState,
}

const StoreContext = createContext(initialState)

const actions = {
  ...userActions,
}

const Reducer = (state, action) => {
  const act = actions[action.type]
  const value = act(state, action)
  return value ? { ...value } : { ...state }
}

export const StoreProvider = ({ children }) => {
  const [state, dispatch] = useReducer(Reducer, initialState)
  const { Provider } = StoreContext
  return <Provider value={{ state, dispatch }}>{children}</Provider>
}

StoreProvider.propTypes = {
  children: PropTypes.element,
}

export const useStore = () => {
  const { state, dispatch } = useContext(StoreContext)
  return { state, dispatch }
}
