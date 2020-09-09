import React, { useState, useImperativeHandle } from 'react'
import { Modal, Form, Input, InputNumber, notification } from 'antd'
import { useForm } from 'antd/lib/form/Form'
import { forwardRef } from 'react'

import { editDictionaryDetail, addDictionaryDetail } from '@/apis/system'

const DictionaryForm = ({ onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [formIns] = useForm()
  const [formData, setFormData] = useState({})

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      const submitData = { ...formData, ...values }
      submitData.id ? await editDictionaryDetail(submitData) : await addDictionaryDetail(submitData)
      notification.success({ message: '操作成功' })
      setVisible(false)
      onSuccess()
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = {}) {
    const formData = {
      dictionaryId: data.dictionaryId,
      id: data.id || '',
      label: data.label || '',
      value: data.value || '',
      sort: data.sort ? Number(data.sort) : 1
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
      title={formData.id ? '编辑字典详情' : '新建字典详情'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} labelCol={{ span: 4 }} onFinish={handleSubmit}>
        <Form.Item name="label" label="字典标签" rules={[{ required: true }]}>
          <Input autoFocus />
        </Form.Item>
        <Form.Item name="value" label="字典值" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item name="sort" label="排序">
          <InputNumber style={{ width: '100%' }} />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(DictionaryForm)
