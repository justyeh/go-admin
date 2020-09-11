import React, { useState, forwardRef } from 'react'
import { Form, Input, Button, notification } from 'antd'
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
    <Form labelCol={{ span: 6 }} form={formIns} onFinish={handleSubmit}>
      <Form.Item name="address" label="邮件地址" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item name="code" label="授权码" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item name="name" label="发信名称" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item name="smtp" label="发送服务器(SMTP)" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 6 }}>
        <Button type="primary" htmlType="submit" loading={submitLoading}>
          保存
        </Button>
      </Form.Item>
    </Form>
  )
}

export default forwardRef(DeptForm)
