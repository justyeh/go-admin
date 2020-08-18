import { createStore, applyMiddleware } from 'redux'

// import logger from 'redux-logger'
import thunk from 'redux-thunk'
import reducer from './reducer'

// const middleware = process.env.NODE_ENV === 'development' ? [logger, thunk] : [thunk]
const middleware = [thunk]
const store = createStore(reducer, applyMiddleware(...middleware))

export default store
