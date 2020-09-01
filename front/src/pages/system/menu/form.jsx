import React, { useState, useImperativeHandle } from 'react'
import { Modal, Form, Input, InputNumber, TreeSelect } from 'antd'
import { useForm } from 'antd/lib/form/Form'
import { forwardRef } from 'react'

const MenuForm = ({ onSuccess }, ref) => {
  const [visible, setVisible] = useState(false)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [menuData] = useState([{ id: '0', name: '顶级菜单' }])
  const [initialValues, setInitialValues] = useState({})

  const [formIns] = useForm()

  const handleSubmit = async (values) => {
    setSubmitLoading(true)
    try {
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = {}) {
    formIns.resetFields()
    setInitialValues({
      icon: data.icon || '',
      name: data.icon || '',
      url: data.icon || '',
      component: data.icon || '',
      sort: data.sort ? Number(data.sort) : 1,
      metaDate: data.icon || '',
      pid: undefined
    })
    setVisible(true)
  }

  const validateComponent = (rule, value) => {
    const url = formIns.getFieldValue('url')
    if (url && !url.startsWith('http') && value.trim().length === 0) {
      return Promise.reject('请填写组件路径')
    }
    return Promise.resolve()
  }

  useImperativeHandle(ref, () => {
    return { init }
  })

  return (
    <Modal
      title={initialValues.id ? '编辑菜单' : '新建菜单'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={formIns.submit}
    >
      <Form form={formIns} labelCol={{ span: 4 }} initialValues={initialValues} onFinish={handleSubmit}>
        <Form.Item name="icon" label="图标">
          <Input placeholder="请输入对应的图标antd icon名称，例如：MenuOutlined" />
        </Form.Item>
        <Form.Item name="name" label="名称" rules={[{ required: true, message: '请填写菜单名称!' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="url" label="链接" rules={[{ required: true, message: '请填写菜单链接!' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="component" label="组件路径" rules={[{ validator: validateComponent }]}>
          <Input />
        </Form.Item>
        <Form.Item name="sort" label="排序">
          <InputNumber style={{ width: '100%' }} />
        </Form.Item>
        <Form.Item name="metaData" label="菜单数据">
          <Input.TextArea placeholder="格式：key=val，多个使用请使用#分割" />
        </Form.Item>
        <Form.Item name="pid" label="父级菜单" rules={[{ required: true, message: '请选择父级菜单!' }]}>
          <TreeSelect treeData={menuData} treeDataSimpleMode={{ id: 'name', pId: 'pid' }} />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default forwardRef(MenuForm)
