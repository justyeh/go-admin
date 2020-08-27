import React from 'react'
import Header from '@/components/Header'
import Sider from '@/components/Sider'
import Breadcrumb from '@/components/Breadcrumb'

export default ({ children }) => {
  return (
    <div className="layout layout-management">
      <Header />
      <div className="container">
        <Sider />
        <div className="page">
          <Breadcrumb />
          {children}
        </div>
      </div>
    </div>
  )
}
