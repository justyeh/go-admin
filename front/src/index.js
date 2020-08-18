import React from 'react'
import ReactDOM from 'react-dom'
import App from '@/App'
import * as serviceWorker from '@/serviceWorker'
import { notification } from 'antd'

import '@/assets/style/antd.scss';
import '@/assets/style/global.scss'

notification.config({
  duration: 2
})

ReactDOM.render(<App />, document.getElementById('root'))

serviceWorker.unregister()
