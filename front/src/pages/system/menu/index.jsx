import React, { useState, Fragment } from 'react'
import { Table, Button, Input } from 'antd'
import { useHistory } from 'react-router-dom'
import { getQueryVariable } from '@/utils/index'

import { menuTree } from '@/apis/system'
import { useMount } from 'react-use'

const tableColumns = [
  {
    title: '菜单名称',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: '链接地址',
    dataIndex: 'url',
    key: 'url'
  },
  {
    title: '组件路径',
    dataIndex: 'path',
    key: 'path'
  },
  {
    title: '排序',
    dataIndex: 'sort',
    key: 'sort'
  },
  {
    title: '图标',
    dataIndex: 'icon',
    key: 'icon'
  },
  {
    title: '菜单数据',
    dataIndex: 'data',
    key: 'data'
  },
  {
    title: '操作',
    dataIndex: 'handle',
    key: 'handle',
    render() {
      return (
        <Fragment>
          <Button>编辑</Button>
          <Button>删除</Button>
        </Fragment>
      )
    }
  }
]

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [tableData, setTableData] = useState([])

  function handleSearch(e) {
    history.push('/system/menu?keyword=' + e)
  }

  async function getDataList() {
    try {
      const { list = [] } = await menuTree({ keyword })
      setTableData(list)
    } catch (error) {}
  }

  function handleAdd() {}

  useMount(getDataList)

  return (
    <div className="menu-page">
      <div className="page-filter-box">
        <Input.Search placeholder="菜单名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加菜单
        </Button>
      </div>
      <Table columns={tableColumns} dataSource={tableData} />
    </div>
  )
}
