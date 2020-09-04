import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification, Switch, Row, Col, Tabs, Spin, Tree } from 'antd'
import { getQueryVariable, bindPage, convertAntdNodeData } from '@/utils'
import RoleForm from './form'
import Pagination from '@/components/Pagination'
import qs from 'qs'

import { roleList, delRole, updateRoleStatus, menuTree, permissionTree } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [page, setPage] = useState(bindPage)
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const [selectedRowKeys, setSelectedRowKeys] = useState([])

  const [menuData, setMenuData] = useState([])
  const [permissionData, setPermissionData] = useState([])
  const [relyLoading, setRelyLoading] = useState(false)

  const handleSearch = (e) => {
    history.push('/system/role?keyword=' + e)
  }

  const handlePageChange = (current, size, isReplace = false) => {
    history[!!isReplace ? 'replace' : 'push'](
      '/system/job?' +
        qs.stringify({
          keyword: keyword,
          current: current,
          size: size
        })
    )
  }

  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [], total = 0 } = await roleList({ keyword, current: page.current, size: page.size })
      setTableData(list)
      setPage((val) => ({ ...val, total: total }))
    } catch (error) {}
    setTableLoading(false)
  }

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
      menuData = convertAntdNodeData(menuData)
      permissionData = convertAntdNodeData(permissionData)
      setMenuData(menuData)
      setPermissionData(permissionData)
      sessionStorage.setItem(
        'RelyDataCache',
        JSON.stringify({
          exp: Date.now() + 1000 * 5, // 数据缓存5s
          menuData,
          permissionData
        })
      )
    } catch (error) {}
    setRelyLoading(false)
  }

  const handleAdd = () => {
    formRef.current.init()
  }

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

  const handleEdit = (data) => {
    formRef.current.init(data)
  }

  const handleChangeStatus = async ({ id }, status) => {
    setTableLoading(true)
    try {
      await updateRoleStatus({ id, status: status ? 'active' : 'ban' })
      getTableData()
    } catch (error) {
      setTableLoading(false)
    }
  }

  const handleSelectionChange = (keys) => {
    setSelectedRowKeys(keys)
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
            <Tabs defaultActiveKey="menu" tabBarExtraContent={<Button type="primary">保存</Button>}>
              <Tabs.TabPane key="menu" tab="菜单分布">
                <Tree checkable disabled={!selectedRowKeys[0]} treeData={menuData} />
              </Tabs.TabPane>
              <Tabs.TabPane key="permission" tab="权限分配">
                <Tree checkable disabled={!selectedRowKeys[0]} treeData={permissionData} />
              </Tabs.TabPane>
            </Tabs>
          </Spin>
        </Col>
      </Row>

      <RoleForm ref={formRef} onSuccess={getTableData} />
    </div>
  )
}
