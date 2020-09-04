import React, { useState, Fragment, useRef } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification } from 'antd'
import { getQueryVariable, dateFormat } from '@/utils'
import DeptForm from './form'

import { deptTree, delDept } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const handleSearch = (e) => {
    history.push('/system/dept?keyword=' + e)
  }

  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [] } = await deptTree({ keyword })
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
          await delDept(id)
          notification.success({ message: '操作成功' })
          getTableData()
        } catch (error) {}
        setTableLoading(false)
      }
    })
  }

  const handleEdit = (data) => {
    formRef.current.init(data)
  }

  useMount(getTableData)

  return (
    <div className="dept-page">
      <div className="page-filter-box">
        <Input.Search placeholder="部门名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加部门
        </Button>
      </div>
      <Table loading={tableLoading} pagination={false} size="small" rowKey="id" dataSource={tableData}>
        <Table.Column dataIndex="name" title="部门名称" />
        <Table.Column dataIndex="sort" title="排序" align="center" />
        <Table.Column title="创建时间" align="center" render={(row) => dateFormat(row.createAt)} />
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
      <DeptForm ref={formRef} deptData={tableData} onSuccess={getTableData} />
    </div>
  )
}
