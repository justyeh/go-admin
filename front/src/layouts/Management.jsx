import React from 'react'
import Header from '../components/Header'
import Sider from '../components/Sider'

export default ({ children }) => {
  return (
    <div className="layout layout-management">
      <Header />
      <Sider />
    </div>
  )
}
