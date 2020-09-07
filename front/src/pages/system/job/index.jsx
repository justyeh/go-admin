import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, notification } from 'antd'
import { getQueryVariable, bindPage, dateFormat } from '@/utils'
import JobForm from './form'
import Pagination from '@/components/Pagination'
import qs from 'qs'

import { jobList, delJob } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [page, setPage] = useState(bindPage)
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const handleSearch = (e) => {
    doSearch(e, { current: 1, size: page.size })
  }

  const handlePageChange = (current, size, isReplace = false) => {
    doSearch(keyword, { current, size }, isReplace)
  }

  const doSearch = (keyword = '', page = { current: 1, size: 10 }, isReplace = false) => {
    history[!!isReplace ? 'replace' : 'push']('/system/job?' + qs.stringify({ keyword, ...page }))
  }

  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [], total = 0 } = await jobList({ keyword, current: page.current, size: page.size })
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
          await delJob(id)
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

  useMount(getTableData)

  return (
    <div className="job-page">
      <div className="page-filter-box">
        <Input.Search placeholder="岗位名称" enterButton allowClear defaultValue={keyword} onSearch={handleSearch} />
        <Button type="primary" onClick={handleAdd}>
          添加岗位
        </Button>
      </div>
      <Table loading={tableLoading} pagination={false} size="small" rowKey="id" dataSource={tableData}>
        <Table.Column dataIndex="name" title="岗位名称" />
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

      <Pagination page={page} onChange={handlePageChange} />
      <JobForm ref={formRef} onSuccess={getTableData} />
    </div>
  )
}
