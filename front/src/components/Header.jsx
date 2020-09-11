import React from 'react'
import { NavLink } from 'react-router-dom'

export default function Header() {
  return (
    <div className="header">
      <div className="title">G CMS</div>
      <div className="nav">
        <NavLink to="/system">系统管理</NavLink>
        <NavLink to="/application">应用中心</NavLink>
        <NavLink to="/website">网站设置</NavLink>
      </div>
    </div>
  )
}
