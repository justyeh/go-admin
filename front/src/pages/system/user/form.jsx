import React, { useState, useImperativeHandle, forwardRef } from 'react'
import { Modal, Form, Input, Switch, notification, Select, TreeSelect } from 'antd'
import { useForm } from 'antd/lib/form/Form'
import { convertAntdNodeData } from '@/utils/'

import { editUser, addUser } from '@/apis/system'
import { useEffect } from 'react'

const UserForm = ({ rely, onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [formIns] = useForm()
  const [formData, setFormData] = useState({})

  const [deptData, setDeptData] = useState([])
  const [jobData, setJobData] = useState([])
  const [roleData, setRoleData] = useState([])

  useEffect(() => {
    rely.deptData && setDeptData(rely.deptData)
    rely.jobData && setJobData(rely.jobData)
    rely.roleData && setRoleData(rely.roleData)
  }, [rely])

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      const submitData = { ...formData, ...values }
      submitData.status = submitData.status ? 'active' : 'ban'
      submitData.id ? await editUser(submitData) : await addUser(submitData)
      notification.success({ message: '操作成功' })
      setVisible(false)
      onSuccess()
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = { status: 'active', job: {}, dept: {}, roleList: [] }) {
    const formData = {
      id: data.id || '',
      account: data.account || '',
      nickname: data.nickname || '',
      status: data.status === 'active',
      phone: data.phone || '',
      email: data.email || '',
      jobId: data.job.id || '',
      deptId: data.dept.id || '',
      roleIds: data.roleList.map((item) => item.id)
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
      title={formData.id ? '编辑用户' : '新建用户'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={600}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} className="user-form" onFinish={handleSubmit}>
        <Form.Item
          name="account"
          label="账号"
          rules={[{ required: true }, { pattern: /^[a-zA-Z0-9_]*$/, message: '仅允许字母、数字、下划线' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item name="nickname" label="昵称" rules={[{ required: true }]}>
          <Input />
        </Form.Item>
        <Form.Item
          name="phone"
          label="手机"
          rules={[{ required: true }, { pattern: /^[a-zA-Z0-9_]*$/, message: '手机号格式错误' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item name="email" label="邮箱" rules={[{ required: true }, { type: 'email' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="deptId" label="部门" rules={[{ required: true }]}>
          <TreeSelect treeData={deptData} />
        </Form.Item>
        <Form.Item name="jobId" label="岗位" rules={[{ required: true }]}>
          <Select options={jobData} />
        </Form.Item>
        <Form.Item name="status" label="状态" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="停用" />
        </Form.Item>
        <Form.Item name="roleIds" label="角色" className="full" rules={[{ required: true }]}>
          <Select mode="multiple" options={roleData} />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(UserForm)
