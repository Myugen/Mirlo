import React, { createContext, useReducer } from 'react'
import PropTypes from 'prop-types'

const initialState = {
  user: undefined,
}

const GlobalStoreContext = createContext(initialState)

const Reducer = (state, action) => {
  const actions = {
    SET_USER: () => {
      const { user } = action.payload
      return {
        ...state,
        user,
      }
    },
    GET_USER: () => {
      const { user } = state
      return user
    },
  }

  return actions[action.type] || state
}

const GlobalStore = ({ children }) => {
  const [state, dispatch] = useReducer(Reducer, initialState)
  const { Provider } = GlobalStoreContext
  return <Provider value={{ state, dispatch }}>{children}</Provider>
}

export { GlobalStore, GlobalStoreContext }

GlobalStore.propTypes = {
  children: PropTypes.element,
}
