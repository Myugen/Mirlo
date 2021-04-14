export const userInitialState = {
  user: undefined,
}

export const userActions = {
  SET_USER: (state, { payload = {} }) => {
    const { user } = payload
    return {
      ...state,
      user,
    }
  },
  GET_USER: (state = {}) => {
    const { user } = state
    return user
  },
}
