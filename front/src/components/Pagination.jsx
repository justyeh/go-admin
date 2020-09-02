import React from 'react'
import { Pagination } from 'antd'
export default ({ page, onChange }) => {
  function onShowSizeChange(pageCurrent, pageSize) {
    onChange(1, pageSize)
  }
  function onPageCurrentChange(pageCurrent, pageSize) {
    onChange(pageCurrent, pageSize)
  }
  return page.total === 0 ? null : (
    <div className="pagination-box">
      <Pagination
        size="small"
        current={page.current}
        pageSize={page.size}
        total={page.total}
        onChange={onPageCurrentChange}
        showSizeChanger
        onShowSizeChange={onShowSizeChange}
        defaultCurrent={3}
        showTotal={(total) => `共计${total}条数据`}
      />
    </div>
  )
}
