import React, { useState, useImperativeHandle } from 'react'
import { Modal, Form, Input, InputNumber, notification } from 'antd'
import { useForm } from 'antd/lib/form/Form'
import { forwardRef } from 'react'

import { editJob, addJob } from '@/apis/system'

const JobForm = ({ onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [formIns] = useForm()
  const [formData, setFormData] = useState({})

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      const submitData = { ...formData, ...values }
      submitData.id ? await editJob(submitData) : await addJob(submitData)
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
      sort: data.sort ? Number(data.sort) : 1,
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
      title={formData.id ? '编辑岗位' : '新建岗位'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} labelCol={{ span: 4 }} onFinish={handleSubmit}>
        <Form.Item name="name" label="名称" rules={[{ required: true, message: '请填写岗位名称!' }]}>
          <Input autoFocus />
        </Form.Item>
        <Form.Item name="sort" label="排序">
          <InputNumber style={{ width: '100%' }} />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(JobForm)
