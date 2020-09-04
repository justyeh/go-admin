import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification } from 'antd'
import AntIcon from '@/components/AntIcon'
import { getQueryVariable } from '@/utils/index'
import MenuForm from './form'

import { menuTree, delMenu } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const handleSearch = (e) => {
    history.push('/system/menu?keyword=' + e)
  }

  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [] } = await menuTree({ keyword })
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
          await delMenu(id)
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
    <div className="menu-page">
      <div className="page-filter-box">
        <Input.Search placeholder="菜单名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加菜单
        </Button>
      </div>
      <Table loading={tableLoading} pagination={false} size="small" rowKey="id" dataSource={tableData}>
        <Table.Column dataIndex="name" title="菜单名称" />
        <Table.Column title="icon" align="center" render={(row) => (row.icon ? <AntIcon name={row.icon} /> : '--')} />
        <Table.Column title="url" align="center" render={(row) => <span>{row.url || '--'}</span>} />
        <Table.Column title="component" align="center" render={(row) => <span>{row.component || '--'}</span>} />
        <Table.Column dataIndex="sort" title="排序" align="center" />
        <Table.Column title="菜单数据" align="center" render={(row) => <span>{row.metaDate || '--'}</span>} />
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
      <MenuForm ref={formRef} menuData={tableData} onSuccess={getTableData} />
    </div>
  )
}
