import React from 'react'
import { Tabs } from 'antd'
import Basic from './components/Basic'
import Code from './components/Code'
import Email from './components/Email'

import './style.scss'

const tabs = [
  {
    label: '基本设置',
    key: '',
    component: <Basic />
  },
  {
    label: '第三方代码',
    key: 'code',
    component: <Code />
  },
  {
    label: '邮箱配置',
    key: 'email',
    component: <Email />
  }
]

export default function NotFount() {
  return (
    <div className="website-page">
      <Tabs>
        {tabs.map((item) => (
          <Tabs.TabPane key={item.key} tab={item.label}>
            {item.component}
          </Tabs.TabPane>
        ))}
      </Tabs>
    </div>
  )
}
