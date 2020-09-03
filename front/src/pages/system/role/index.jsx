import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification, Tag } from 'antd'
import { getQueryVariable } from '@/utils'
import RoleForm from './form'

import { roleList, delRole } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const handleSearch = (e) => {
    history.push('/system/role?keyword=' + e)
  }

  const getDataList = async () => {
    setTableLoading(true)
    try {
      const { list = [] } = await roleList({ keyword })
      setTableData(list)
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
          getDataList()
        } catch (error) {}
        setTableLoading(false)
      }
    })
  }

  const handleEdit = (data) => {
    formRef.current.init(data)
  }

  useMount(getDataList)

  return (
    <div className="role-page">
      <div className="page-filter-box">
        <Input.Search placeholder="角色名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加角色
        </Button>
      </div>
      <Table loading={tableLoading} pagination={false} size="small" rowKey="id" dataSource={tableData}>
        <Table.Column dataIndex="name" title="角色名称" />
        <Table.Column dataIndex="code" title="角色CODE" align="center" />
        <Table.Column
          title="角色CODE"
          align="center"
          render={(row) => <Tag>{row.status === 'Active' ? '启用' : '停用'}</Tag>}
        />
        <Table.Column dataIndex="sort" title="排序" align="center" />
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
      <RoleForm ref={formRef} onSuccess={getDataList} />
    </div>
  )
}
