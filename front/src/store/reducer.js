const initialState = {
  menus: [],
  user: {}
}

const rootReducer = (state = initialState, action) => {
  switch (action.type) {
    case 'logout':
      return { ...state, user: {} }
    case 'updateUser':
      state.user = { ...state.user, ...action.user }
      return { ...state }
    default:
      return state
  }
}

export default rootReducer
