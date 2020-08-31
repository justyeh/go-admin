import React from 'react'
import { useState } from 'react'
import { Modal, Form, Input, InputNumber, TreeSelect } from 'antd'
import { useForm } from 'antd/lib/form/Form'

export default () => {
  const [visible, setVisible] = useState(true)
  const [submitLoading, setSubmitLoading] = useState(false)
  const [menuData,setMenuData] = useState([{id:"0",name:"顶级菜单"}])
  const [form, setForm] = useState({})

  const [formIns] = useForm()

  const handleSubmit = async (values) => {
    formIns.validateFields()
    setSubmitLoading(true)
    try {
      console.log(123)
    } catch (error) {}
    setSubmitLoading(false)
  }

  function init(data = {}) {
    setForm({ ...data })
    setVisible(true)
  }

  return (
    <Modal
      title={form.id ? '编辑菜单' : '新建菜单'}
      maskClosable={false}
      keyboard={false}
      visible={visible}
      width={500}
      confirmLoading={submitLoading}
      onCancel={() => setVisible(false)}
      onOk={handleSubmit}
    >
      <Form form={formIns} name="menu" labelCol={{ span: 4 }}>
        <Form.Item name="icon" label="图标">
          <Input placeholder="请输入对应的图标antd icon名称，例如：MenuOutlined" />
        </Form.Item>
        <Form.Item name="name" label="名称" rules={[{ required: true, message: '请填写菜单名称!' }]}>
          <Input />
        </Form.Item>
        <Form.Item name="url" label="链接" rules={[{ required: true, message: '请填写菜单链接!' }]}>
          <Input />
        </Form.Item>
        <Form.Item
          name="component"
          label="组件路径"
          rules={[
            ({ getFieldValue }) => ({
              validator(rule, value) {
                if (getFieldValue('url').startsWith('http') && value.trim().length === 0) {
                  return Promise.reject('请填写组件路径')
                }
                return Promise.resolve()
              }
            })
          ]}
        >
          <Input />
        </Form.Item>
        <Form.Item name="sort" label="排序">
          <InputNumber style={{ width: '100%' }} />
        </Form.Item>
        <Form.Item name="metaData" label="菜单数据">
          <Input />
        </Form.Item>
        <Form.Item name="metaData" label="父级菜单" rules={[{ required: true, message: '请选择父级菜单!' }]}>
          <TreeSelect treeData={menuData} treeDataSimpleMode={{id:"name",pId:"pid"}}/>
        </Form.Item>
      </Form>
    </Modal>
  )
}
