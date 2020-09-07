import React, { useRef, useState, Fragment } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Table, Button, Input, Modal, Row, Col, Card, notification } from 'antd'
import { getQueryVariable, bindPage } from '@/utils'
import DictionaryForm from './form'
import DictionaryDetailForm from './detail-form'
import Pagination from '@/components/Pagination'
import qs from 'qs'
import './style.scss'

import { dictionaryList, delDictionary, dictionaryDetailList, delDictionaryDetail } from '@/apis/system'

export default () => {
  const history = useHistory()

  const [keyword] = useState(getQueryVariable('keyword'))
  const [page, setPage] = useState(bindPage)
  const [tableData, setTableData] = useState([])
  const [tableLoading, setTableLoading] = useState(false)
  const formRef = useRef()

  const [selectedRowKeys, setSelectedRowKeys] = useState([])

  const [detailTableData, setDetailTableData] = useState([])
  const [detailTableLoading, setDetailTableLoading] = useState(false)
  const detailFormRef = useRef()

  const handleSearch = (e) => {
    doSearch(e, { current: 1, size: page.size })
  }

  const handlePageChange = (current, size, isReplace = false) => {
    doSearch(keyword, { current, size }, isReplace)
  }

  const doSearch = (keyword = '', page = { current: 1, size: 10 }, isReplace = false) => {
    history[!!isReplace ? 'replace' : 'push']('/system/dictionary?' + qs.stringify({ keyword, ...page }))
  }

  /* 获取字典列表 */
  const getTableData = async () => {
    setTableLoading(true)
    try {
      const { list = [], total = 0 } = await dictionaryList({ keyword, current: page.current, size: page.size })
      setTableData(list)
      setPage((val) => ({ ...val, total: total }))
    } catch (error) {}
    setTableLoading(false)
  }

  /* 字典新增 */
  const handleAdd = () => {
    formRef.current.init()
  }

  /* 字典删除 */
  const handleDelete = ({ id }) => {
    Modal.confirm({
      title: '确认删除该数据吗？',
      onOk: async () => {
        setTableLoading(true)
        try {
          await delDictionary(id)
          notification.success({ message: '操作成功' })
          const current = tableData.length === 1 ? --page.current : page.current
          handlePageChange(current, page.size, 1)
        } catch (error) {}
        setTableLoading(false)
      }
    })
  }

  /* 字典编辑 */
  const handleEdit = (data) => {
    formRef.current.init(data)
  }

  /* 获取字典详情 */
  const getDictionaryDetailList = async (dictionaryId) => {
    if (!dictionaryId) {
      return
    }
    setDetailTableLoading(true)
    try {
      const { list = [] } = await dictionaryDetailList({ dictionaryId })
      setDetailTableData(list)
    } catch (error) {}
    setDetailTableLoading(false)
  }

  /* 字典详情新增 */
  const handleDetailAdd = () => {
    detailFormRef.current.init({ dictionaryId: selectedRowKeys[0] })
  }

  /* 字典详情删除 */
  const handleDetailDelete = ({ id }) => {
    Modal.confirm({
      title: '确认删除该数据吗？',
      onOk: async () => {
        setDetailTableLoading(true)
        try {
          await delDictionaryDetail(id)
          notification.success({ message: '操作成功' })
          getDictionaryDetailList(selectedRowKeys[0])
        } catch (error) {}
        setDetailTableLoading(false)
      }
    })
  }

  /* 字典详情编辑 */
  const handleDetailEdit = (data) => {
    detailFormRef.current.init({ ...data, dictionaryId: selectedRowKeys[0] })
  }

  const handleSelectionChange = (keys) => {
    setSelectedRowKeys(keys)
    getDictionaryDetailList(keys[0])
  }

  useMount(getTableData)

  return (
    <div className="dictionary-page">
      <Row gutter={20}>
        <Col span={14}>
          <Card title="字典列表">
            <div className="page-filter-box" style={{ paddingTop: 0 }}>
              <Input.Search
                placeholder="字典名称"
                enterButton
                allowClear
                defaultValue={keyword}
                onSearch={handleSearch}
              />
              <Button type="primary" onClick={handleAdd}>
                添加字典
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
              <Table.Column dataIndex="name" title="字典名称" />
              <Table.Column title="字典描述" align="center" render={(row) => row.description || '--'} />
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
          </Card>
        </Col>
        <Col span={10}>
          <Card title="字典详情">
            {selectedRowKeys[0] ? (
              <Fragment>
                <div className="page-filter-box" style={{ paddingTop: 0 }}>
                  <Button type="primary" onClick={handleDetailAdd}>
                    添加字典详情
                  </Button>
                </div>
                <Table
                  loading={detailTableLoading}
                  pagination={false}
                  size="small"
                  rowKey="id"
                  dataSource={detailTableData}
                >
                  <Table.Column dataIndex="label" title="字典标签" align="center" />
                  <Table.Column dataIndex="value" title="字典值" align="center" />
                  <Table.Column dataIndex="sort" title="字典值" align="center" />
                  <Table.Column
                    title="操作"
                    align="center"
                    render={(row) => (
                      <Fragment>
                        <Button type="link" onClick={() => handleDetailDelete(row)}>
                          删除
                        </Button>
                        <Button type="link" onClick={() => handleDetailEdit(row)}>
                          编辑
                        </Button>
                      </Fragment>
                    )}
                  />
                </Table>
              </Fragment>
            ) : (
              <div className="tip">选中字典查看详情</div>
            )}
          </Card>
        </Col>
      </Row>

      <DictionaryForm ref={formRef} onSuccess={getTableData} />
      <DictionaryDetailForm ref={detailFormRef} onSuccess={() => getDictionaryDetailList(selectedRowKeys[0])} />
    </div>
  )
}
