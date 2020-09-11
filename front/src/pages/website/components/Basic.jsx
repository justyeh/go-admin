import React, { useState, forwardRef } from 'react'
import { Form, Input, Button, Upload, notification } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import { useForm } from 'antd/lib/form/Form'

const DeptForm = () => {
  const [submitLoading, setSubmitLoading] = useState(false)
  const [formIns] = useForm()

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      notification.success({ message: '保存成功' })
    } catch (error) {}
    setSubmitLoading(false)
  }

  return (
    <Form labelCol={{ span: 4 }} form={formIns} onFinish={handleSubmit}>
      <Form.Item name="name" label="网站名称" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item name="logo" label="站点logo">
        <Upload name="avatar" listType="picture-card" showUploadList={false}>
          <div>
            <PlusOutlined />
            <div style={{ marginTop: 8 }}>Upload</div>
          </div>
        </Upload>
      </Form.Item>
      <Form.Item name="domain" label="站点域名">
        <Input />
      </Form.Item>
      <Form.Item name="logo" label="版权信息">
        <Input />
      </Form.Item>
      <Form.Item name="logo" label="备案信息">
        <Input />
      </Form.Item>
      <Form.Item name="logo" label="主营行业">
        <Input />
      </Form.Item>
      <Form.Item name="logo" label="主营产品">
        <Input />
      </Form.Item>
      <Form.Item name="keyword" label="全站关键词">
        <Input />
      </Form.Item>
      <Form.Item name="description" label="网站介绍">
        <Input.TextArea />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 4 }}>
        <Button type="primary" htmlType="submit" loading={submitLoading}>
          保存
        </Button>
      </Form.Item>
    </Form>
  )
}

export default forwardRef(DeptForm)
