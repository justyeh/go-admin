import React, { useState, Fragment } from 'react'
import { Table, Button, Input } from 'antd'
import { useHistory } from 'react-router-dom'
import { getQueryVariable } from '@/utils/index'
import MenuForm from './form'

import { menuTree } from '@/apis/system'
import { useMount } from 'react-use'
import { useRef } from 'react'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [tableData, setTableData] = useState([])
  const formRef = useRef()

  const handleSearch = (e) => {
    history.push('/system/menu?keyword=' + e)
  }

  const getDataList = async () => {
    try {
      const { list = [] } = await menuTree({ keyword })
      setTableData(list)
    } catch (error) {}
  }

  const handleAdd = () => {
    formRef.current.init()
  }

  const handleDelete = ({ id }) => {}

  const handleEdit = (data) => {
    formRef.current.init(data)
  }

  useMount(getDataList)

  return (
    <div className="menu-page">
      <div className="page-filter-box">
        <Input.Search placeholder="菜单名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加菜单
        </Button>
      </div>
      <Table dataSource={tableData}>
        <Table.Column key="icon" dataIndex="icon" title="icon" align="center" />
        <Table.Column key="name" dataIndex="name" title="菜单名称" align="center" />
        <Table.Column key="url" dataIndex="url" title="链接" align="center" />
        <Table.Column key="component" dataIndex="component" title="组件路径" align="center" />
        <Table.Column key="sort" dataIndex="sort" title="排序" align="center" />
        <Table.Column key="metaData" dataIndex="metaData" title="菜单数据" align="center" />
        <Table.Column
          key="handle"
          title="操作"
          align="center"
          render={() => (
            <Fragment>
              <Button type="link" onClick={handleDelete}>
                删除
              </Button>
              <Button type="link" onClick={handleEdit}>
                编辑
              </Button>
            </Fragment>
          )}
        />
      </Table>
      <MenuForm ref={formRef} menuData={tableData} onSuccess="getDataList" />
    </div>
  )
}
