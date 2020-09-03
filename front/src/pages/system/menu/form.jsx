import React, { useState, useImperativeHandle, forwardRef, useEffect } from 'react'
import { Modal, Form, Input, InputNumber, TreeSelect, notification } from 'antd'
import { useForm } from 'antd/lib/form/Form'

import { editMenu, addMenu } from '@/apis/system'

const { TreeNode } = TreeSelect

function renderTreeNodes(data = []) {
  return data.map((item) => {
    return (
      <TreeNode key={item.id} title={item.name} value={item.id}>
        {item.children && item.children.length > 0 && renderTreeNodes(item.children)}
      </TreeNode>
    )
  })
}

const MenuForm = ({ menuData, onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [treeData, setTreeData] = useState([{ id: '0', name: '顶级菜单' }])
  const [formIns] = useForm()
  const [formData, setFormData] = useState({})

  useEffect(() => {
    setTreeData([{ id: '0', name: '顶级菜单', children: menuData }])
  }, [menuData])

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
      const submitData = { ...formData, ...values }
      submitData.id ? await editMenu(submitData) : await addMenu(submitData)
      notification.success({ message: '操作成功' })
      setVisible(false)
      onSuccess()
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = {}) {
    const formData = {
      id: data.id || '',
      icon: data.icon || '',
      name: data.name || '',
      url: data.url || '',
      component: data.component || '',
      sort: data.sort ? Number(data.sort) : 1,
      metaDate: data.metaDate || '',
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
      title={formData.id ? '编辑菜单' : '新建菜单'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} labelCol={{ span: 4 }} onFinish={handleSubmit}>
        <Form.Item name="icon" label="图标">
          <Input placeholder="请输入对应的图标antd icon名称，例如：MenuOutlined" />
        </Form.Item>
        <Form.Item name="name" label="名称" rules={[{ required: true, message: '请填写菜单名称!' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="url" label="链接">
          <Input />
        </Form.Item>
        <Form.Item name="component" label="组件路径">
          <Input />
        </Form.Item>
        <Form.Item name="sort" label="排序">
          <InputNumber style={{ width: '100%' }} />
        </Form.Item>
        <Form.Item name="metaData" label="菜单数据">
          <Input.TextArea placeholder="格式：key=val，多个使用请使用#分割" />
        </Form.Item>
        <Form.Item name="pid" label="父级菜单" rules={[{ required: true, message: '请选择父级菜单!' }]}>
          <TreeSelect>{renderTreeNodes(treeData)}</TreeSelect>
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(MenuForm)
