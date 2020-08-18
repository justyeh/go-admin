import React, { memo } from 'react'
import Router from './router/index'

import zhCN from 'antd/es/locale/zh_CN'
import { ConfigProvider } from 'antd'

import { Provider } from 'react-redux'
import store from '@/store/index'

const App = memo(() => {
  return (
    <ConfigProvider locale={zhCN}>
      <Provider store={store}>
        <Router />
      </Provider>
    </ConfigProvider>
  )
})

export default App
