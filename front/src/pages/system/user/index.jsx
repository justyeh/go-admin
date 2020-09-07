import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, Row, Col, Tree, Switch, Spin, notification } from 'antd'
import { DownOutlined } from '@ant-design/icons'
import { getQueryVariable, bindPage, convertAntdNodeData } from '@/utils'
import UserForm from './form'
import Pagination from '@/components/Pagination'
import qs from 'qs'

import './style.scss'

import { userList, delUser, updateUserStatus, deptTree, jobList, roleList } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [page, setPage] = useState(bindPage)
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const [selectDeptId] = useState(getQueryVariable('deptId'))

  const [deptData, setDeptData] = useState([])
  const [jobData, setJobData] = useState([])
  const [roleData, setRoleData] = useState([])
  const [deptLoading, setDeptLoading] = useState(false)

  const handleSearch = (e) => {
    doSearch(e, selectDeptId, { current: 1, size: page.size })
  }

  const handlePageChange = (current, size, isReplace = false) => {
    doSearch(keyword, selectDeptId, { current, size }, isReplace)
  }

  const handleDeptSlected = (selectedKeys) => {
    doSearch(keyword, selectedKeys[0], { current: 1, size: page.size }, true)
  }

  const doSearch = (keyword = '', deptId = '', page = { current: 1, size: 10 }, isReplace = false) => {
    setCache()
    history[!!isReplace ? 'replace' : 'push']('/system/user?' + qs.stringify({ keyword, deptId, ...page }))
  }

  const setCache = () => {
    sessionStorage.setItem(
      'RelyDataCache',
      JSON.stringify({
        exp: Date.now() + 1000 * 5, // 数据缓存5s
        deptData,
        jobData,
        roleData
      })
    )
  }

  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [], total = 0 } = await userList({
        keyword,
        current: page.current,
        size: page.size,
        deptId: selectDeptId
      })
      setTableData(list)
      setPage((val) => ({ ...val, total: total }))
    } catch (error) {}
    setTableLoading(false)
  }

  const getRelyData = async () => {
    try {
      let RelyDataCache = JSON.parse(sessionStorage.getItem('RelyDataCache'))

      if (RelyDataCache && RelyDataCache.exp > Date.now()) {
        const { deptData = [], jobData = [], roleData = [] } = RelyDataCache
        setDeptData(deptData)
        setJobData(jobData)
        setRoleData(roleData)
        return
      }
    } catch (error) {
      console.error(error)
    }

    setDeptLoading(true)
    try {
      let [{ list: deptData = [] }, { list: jobData = [] }, { list: roleData = [] }] = await Promise.all([
        deptTree(),
        jobList(),
        roleList()
      ])
      deptData = convertAntdNodeData({ data: deptData })
      jobData = convertAntdNodeData({ data: jobData })
      roleData = convertAntdNodeData({ data: roleData })
      setDeptData(deptData)
      setJobData(jobData)
      setRoleData(roleData)
    } catch (error) {}
    setDeptLoading(false)
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
          await delUser(id)
          notification.success({ message: '操作成功' })
          const current = tableData.length === 1 ? --page.current : page.current
          handlePageChange(current, page.size, true)
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
      await updateUserStatus({ id, status: status ? 'active' : 'ban' })
      getTableData()
    } catch (error) {
      setTableLoading(false)
    }
  }

  useMount(() => {
    getTableData()
    getRelyData()
  })

  return (
    <div className="user-page">
      <Row gutter={20} style={{ paddingTop: 20 }}>
        <Col span={7}>
          <Spin spinning={deptLoading}>
            <div style={{ minHeight: 200 }}>
              {deptData.length > 0 && (
                <Tree
                  defaultExpandAll
                  selectable
                  switcherIcon={<DownOutlined />}
                  treeData={deptData}
                  defaultSelectedKeys={[selectDeptId]}
                  onSelect={handleDeptSlected}
                />
              )}
            </div>
          </Spin>
        </Col>
        <Col span={17}>
          <div className="page-filter-box" style={{ paddingTop: 0 }}>
            <Input.Search
              placeholder="账号/昵称"
              enterButton
              allowClear
              defaultValue={keyword}
              onSearch={handleSearch}
            />
            <Button type="primary" onClick={handleAdd}>
              添加用户
            </Button>
          </div>
          <Table loading={tableLoading} pagination={false} size="small" rowKey="id" dataSource={tableData}>
            <Table.Column dataIndex="account" title="账号" align="center" />
            <Table.Column dataIndex="nickname" title="昵称" align="center" />
            <Table.Column dataIndex="phone" title="手机" align="center" />
            <Table.Column dataIndex="email" title="邮箱" align="center" />
            <Table.Column title="部门/岗位" align="center" render={(row) => ({})} />
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
            <Table.Column title="角色" align="center" render={(row) => ({})} />
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
      </Row>

      <UserForm ref={formRef} rely={{ deptData, jobData, roleData }} onSuccess={getTableData} />
    </div>
  )
}
