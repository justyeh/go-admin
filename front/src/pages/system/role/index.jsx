import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification, Switch, Row, Col, Tabs, Spin, Tree } from 'antd'
import { getQueryVariable, bindPage, convertAntdNodeData, formatTreeChechkedRelation } from '@/utils'
import RoleForm from './form'
import Pagination from '@/components/Pagination'
import qs from 'qs'

import {
  roleList,
  delRole,
  updateRoleStatus,
  menuTree,
  permissionTree,
  roleMenuList,
  rolePermissionList,
  updateRoleMenu,
  updateRolePermission
} from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [page, setPage] = useState(bindPage)
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const [selectedRowKeys, setSelectedRowKeys] = useState([])
  const [checkedMenuKeys, setCheckedMenuKeys] = useState({ fullChecked: [], halfChecked: [] })
  const [checkedPermissionKeys, setCheckedPermissionKeys] = useState({ fullChecked: [], halfChecked: [] })

  const [menuData, setMenuData] = useState([])
  const [permissionData, setPermissionData] = useState([])
  const [relyLoading, setRelyLoading] = useState(false)

  // 查找
  const handleSearch = (e) => {
    doSearch(e, { current: 1, size: page.size })
  }

  const handlePageChange = (current, size, isReplace = false) => {
    doSearch(keyword, { current, size }, isReplace)
  }

  const doSearch = (keyword = '', page = { current: 1, size: 10 }, isReplace = false) => {
    setCache()
    history[!!isReplace ? 'replace' : 'push']('/system/role?' + qs.stringify({ keyword, ...page }))
  }

  const setCache = () => {
    sessionStorage.setItem(
      'RelyDataCache',
      JSON.stringify({
        exp: Date.now() + 1000 * 5, // 数据缓存5s
        menuData,
        permissionData
      })
    )
  }

  // 获取角色列表
  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [], total = 0 } = await roleList({ keyword, current: page.current, size: page.size })
      setTableData(list)
      setPage((val) => ({ ...val, total: total }))
    } catch (error) {}
    setTableLoading(false)
  }

  // 获取所有菜单+角色
  const getRelyData = async () => {
    try {
      let RelyDataCache = JSON.parse(sessionStorage.getItem('RelyDataCache'))

      if (RelyDataCache && RelyDataCache.exp > Date.now()) {
        const { menuData = [], permissionData = [] } = RelyDataCache
        setMenuData(menuData)
        setPermissionData(permissionData)
        return
      }
    } catch (error) {
      console.error(error)
    }

    setRelyLoading(true)
    try {
      let [{ list: menuData = [] }, { list: permissionData = [] }] = await Promise.all([menuTree(), permissionTree()])
      menuData = convertAntdNodeData({ data: menuData })
      permissionData = convertAntdNodeData({ data: permissionData })
      setMenuData(menuData)
      setPermissionData(permissionData)
    } catch (error) {}
    setRelyLoading(false)
  }

  // 添加角色
  const handleAdd = () => {
    formRef.current.init()
  }

  // 删除角色
  const handleDelete = ({ id }) => {
    Modal.confirm({
      title: '确认删除该数据吗？',
      onOk: async () => {
        setTableLoading(true)
        try {
          await delRole(id)
          notification.success({ message: '操作成功' })
          const current = tableData.length === 1 ? --page.current : page.current
          handlePageChange(current, page.size, 1)
        } catch (error) {}
        setTableLoading(false)
      }
    })
  }

  // 编辑角色
  const handleEdit = (data) => {
    formRef.current.init(data)
  }

  // 修改角色状态
  const handleChangeStatus = async ({ id }, status) => {
    setTableLoading(true)
    try {
      await updateRoleStatus({ id, status: status ? 'active' : 'ban' })
      getTableData()
    } catch (error) {
      setTableLoading(false)
    }
  }

  // 处理表格选中事件
  const handleSelectionChange = async (keys) => {
    setRelyLoading(true)
    setSelectedRowKeys(keys)
    try {
      const [{ ids: menuIds = [] }, { ids: permissionIds = [] }] = await Promise.all([
        roleMenuList(keys[0]),
        rolePermissionList(keys[0])
      ])

      setCheckedMenuKeys(formatTreeChechkedRelation(menuData, menuIds))
      setCheckedPermissionKeys(formatTreeChechkedRelation(permissionData, permissionIds))
    } catch (error) {}
    setRelyLoading(false)
  }

  // 保存角色的菜单和权限
  const handleSaveMenuAndPermission = async () => {
    setRelyLoading(true)
    try {
      await Promise.all([
        updateRoleMenu({
          roleId: selectedRowKeys[0],
          menuIds: [...checkedMenuKeys.fullChecked, ...checkedMenuKeys.halfChecked]
        }),
        updateRolePermission({
          roleId: selectedRowKeys[0],
          permissionIds: [...checkedPermissionKeys.fullChecked, ...checkedPermissionKeys.halfChecked]
        })
      ])
      notification.success({ message: '保存成功' })
    } catch (error) {}
    setRelyLoading(false)
  }

  // 处理Tree选中事件
  const handleTreeCheck = async (checkedKeys, evt, target) => {
    if (target === 'menu') {
      setCheckedMenuKeys({ fullChecked: checkedKeys, halfChecked: evt.halfCheckedKeys })
    }
    if (target === 'permission') {
      setCheckedPermissionKeys({ fullChecked: checkedKeys, halfChecked: evt.halfCheckedKeys })
    }
  }

  useMount(() => {
    getTableData()
    getRelyData()
  })

  return (
    <div className="role-page">
      <Row gutter={20}>
        <Col span={16}>
          <div className="page-filter-box">
            <Input.Search
              placeholder="角色名称"
              enterButton
              allowClear
              defaultValue={keyword}
              onSearch={handleSearch}
            />
            <Button type="primary" onClick={handleAdd}>
              添加角色
            </Button>
          </div>
          <Table
            loading={tableLoading}
            pagination={false}
            size="small"
            rowKey="id"
            dataSource={tableData}
            rowSelection={{ type: 'radio', selectedRowKeys: selectedRowKeys, onChange: handleSelectionChange }}
          >
            <Table.Column dataIndex="name" title="角色名称" align="center" />
            <Table.Column
              title="状态"
              align="center"
              render={(row) => (
                <Switch
                  checkedChildren="启用"
                  unCheckedChildren="停用"
                  checked={row.status === 'active'}
                  onChange={(status) => handleChangeStatus(row, status)}
                />
              )}
            />
            <Table.Column title="备注" render={(row) => row.remark || '--'} />
            <Table.Column
              title="操作"
              align="center"
              render={(row) => (
                <Fragment>
                  <Button type="link" onClick={() => handleDelete(row)}>
                    删除
                  </Button>
                  <Button type="link" onClick={() => handleEdit(row)}>
                    编辑
                  </Button>
                </Fragment>
              )}
            />
          </Table>
          <Pagination page={page} onChange={handlePageChange} />
        </Col>
        <Col span={8}>
          <Spin spinning={relyLoading}>
            <Tabs
              defaultActiveKey="menu"
              tabBarExtraContent={
                <Button type="primary" disabled={!selectedRowKeys[0]} onClick={handleSaveMenuAndPermission}>
                  保存
                </Button>
              }
            >
              <Tabs.TabPane key="menu" tab="菜单分布">
                <Tree
                  checkable
                  selectable={false}
                  disabled={!selectedRowKeys[0]}
                  treeData={menuData}
                  checkedKeys={checkedMenuKeys.fullChecked}
                  onCheck={(checkedKeys, evt) => handleTreeCheck(checkedKeys, evt, 'menu')}
                />
              </Tabs.TabPane>
              <Tabs.TabPane key="permission" tab="权限分配">
                <Tree
                  checkable
                  selectable={false}
                  disabled={!selectedRowKeys[0]}
                  treeData={permissionData}
                  checkedKeys={checkedPermissionKeys.fullChecked}
                  onCheck={(checkedKeys, evt) => handleTreeCheck(checkedKeys, evt, 'permission')}
                />
              </Tabs.TabPane>
            </Tabs>
          </Spin>
        </Col>
      </Row>

      <RoleForm ref={formRef} onSuccess={getTableData} />
    </div>
  )
}
