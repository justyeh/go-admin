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
    <Form labelCol={{ span: 3 }} form={formIns} onFinish={handleSubmit}>
      <Form.Item name="code" label="访问统计">
        <Input.TextArea rows={4} placeholder="例：百度统计、友盟统计" />
      </Form.Item>
      <Form.Item name="other" label="其他">
        <Input.TextArea rows={4} placeholder="例：客服系统、bug监控" />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 3 }}>
        <Button type="primary" htmlType="submit" loading={submitLoading}>
          保存
        </Button>
      </Form.Item>
    </Form>
  )
}

export default forwardRef(DeptForm)
