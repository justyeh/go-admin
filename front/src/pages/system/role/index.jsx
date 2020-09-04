import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification, Switch } from 'antd'
import { getQueryVariable, bindPage } from '@/utils'
import RoleForm from './form'
import Pagination from '@/components/Pagination'
import qs from 'qs'

import { roleList, delRole, updateRoleStatus } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [page, setPage] = useState(bindPage)
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

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

  useMount(getTableData)

  return (
    <div className="role-page">
      <div className="page-filter-box">
        <Input.Search placeholder="角色名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加角色
        </Button>
      </div>
      <Table loading={tableLoading} pagination={false} size="small" rowKey="id" dataSource={tableData}>
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
      <RoleForm ref={formRef} onSuccess={getTableData} />
    </div>
  )
}
