import React, { useState, useImperativeHandle } from 'react'
import { Modal, Form, Input, notification } from 'antd'
import { useForm } from 'antd/lib/form/Form'
import { forwardRef } from 'react'

import { editDictionary, addDictionary } from '@/apis/system'

const DictionaryForm = ({ onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [formIns] = useForm()
  const [formData, setFormData] = useState({})

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      const submitData = { ...formData, ...values }
      submitData.id ? await editDictionary(submitData) : await addDictionary(submitData)
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
      description: data.description || ''
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
      title={formData.id ? '编辑字典' : '新建字典'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} labelCol={{ span: 5 }} onFinish={handleSubmit}>
        <Form.Item
          name="name"
          label="名称(KEY)"
          rules={[{ required: true }, { pattern: /^[a-z_]*$/, message: '仅允许小写字母、下划线' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item name="description" label="描述">
          <Input.TextArea />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(DictionaryForm)
