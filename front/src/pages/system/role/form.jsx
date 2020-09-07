import React, { useState, useImperativeHandle, forwardRef } from 'react'
import { Modal, Form, Input, Switch, notification } from 'antd'
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
      submitData.status = submitData.status ? 'active' : 'ban'
      submitData.id ? await editRole(submitData) : await addRole(submitData)
      notification.success({ message: '操作成功' })
      setVisible(false)
      onSuccess()
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = { status: 'active' }) {
    const formData = {
      id: data.id || '',
      name: data.name || '',
      status: data.status === 'active',
      remark: data.remark || ''
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
          <Input autoFocus />
        </Form.Item>
        <Form.Item name="status" label="状态" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="停用" />
        </Form.Item>
        <Form.Item name="remark" label="备注">
          <Input.TextArea />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(RoleForm)
