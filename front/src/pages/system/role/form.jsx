import React, { useState, useImperativeHandle, forwardRef } from 'react'
import { Modal, Form, Input, InputNumber, Switch, notification } from 'antd'
import { useForm } from 'antd/lib/form/Form'

import { editRole, addRole } from '@/apis/system'

const RoleForm = ({ onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [formIns] = useForm()
  const [formData, setFormData] = useState({})

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      const submitData = { ...formData, ...values }
      submitData.status = submitData.status ? 'Active' : 'Ban'
      submitData.id ? await editRole(submitData) : await addRole(submitData)
      notification.success({ message: '操作成功' })
      setVisible(false)
      onSuccess()
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = {}) {
    const formData = {
      id: data.id || '',
      name: data.name || '',
      code: data.code || '',
      status: data.status === 'Active',
      sort: data.sort ? Number(data.sort) : 1,
      pid: data.pid || '0'
    }
    setFormData(formData)
    formIns.resetFields()
    formIns.setFieldsValue(formData)
    setVisible(true)
  }

  useImperativeHandle(ref, () => {
    return { init }
  })

  return (
    <Modal
      title={formData.id ? '编辑角色' : '新建角色'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} labelCol={{ span: 4 }} onFinish={handleSubmit}>
        <Form.Item name="name" label="名称" rules={[{ required: true, message: '请填写角色名称!' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="code" label="CODE" rules={[{ required: true, message: '请填写角色CODE!' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="status" label="状态" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="停用" />
        </Form.Item>
        <Form.Item name="sort" label="排序">
          <InputNumber style={{ width: '100%' }} />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(RoleForm)
