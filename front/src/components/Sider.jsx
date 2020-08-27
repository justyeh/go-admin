import React, { useState } from 'react'

import { Menu } from 'antd'
import { SettingOutlined } from '@ant-design/icons'
import { useHistory, useLocation } from 'react-router-dom'

export default () => {
  const history = useHistory()
  const location = useLocation()

  const [defaultSelectedKeys] = useState([location.pathname])

  const handleMenuSelect = ({ key }) => {
    history.push(key)
  }

  return (
    <div className="sider">
      <Menu
        defaultSelectedKeys={defaultSelectedKeys}
        defaultOpenKeys={['/sub2']}
        mode="inline"
        theme="light"
        onSelect={handleMenuSelect}
      >
        <Menu.Item key="/">首页</Menu.Item>
        <Menu.SubMenu key="/sub2" icon={<SettingOutlined />} title="系统管理">
          <Menu.Item key="/system/user">用户管理</Menu.Item>
          <Menu.Item key="/system/role">角色管理</Menu.Item>
          <Menu.Item key="/system/permission">权限管理</Menu.Item>
          <Menu.Item key="/system/menu">菜单管理</Menu.Item>

          <Menu.Item key="/system/dept">部门管理</Menu.Item>
          <Menu.Item key="/system/job">岗位管理</Menu.Item>

          <Menu.Item key="/system/dictionary">字典管理</Menu.Item>
        </Menu.SubMenu>
      </Menu>
    </div>
  )
}
