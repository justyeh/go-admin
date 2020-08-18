import React, { useState, useEffect, useCallback } from 'react'
import { Button, Input, Form } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import './style.scss'

import { getCaptcha } from '@/apis/system'

export default () => {
  const [form, setForm] = useState({ uuid: '', captcha: '' })

  const refreshCaptcha = useCallback(async () => {
    const { uuid, captcha } = await getCaptcha()
    setForm({
      uuid,
      captcha
    })
  }, [])

  useEffect(() => {
    refreshCaptcha()
  }, [])

  const onFinish = (values) => {
    console.log('Success:', values)
  }

  return (
    <div className="login-page">
      <Form size="large" name="login" onFinish={onFinish}>
        <h1>G-CMS</h1>
        <Form.Item name="username" rules={[{ required: true, message: '请输入账号!' }]}>
          <Input placeholder="登录账户" prefix={<UserOutlined />} />
        </Form.Item>

        <Form.Item name="password" rules={[{ required: true, message: '请输入密码!' }]}>
          <Input.Password placeholder="登录密码" prefix={<LockOutlined />} />
        </Form.Item>

        <Form.Item name="captcha" className="captcha-item" rules={[{ required: true, message: '请输入验证码!' }]}>
          <Input placeholder="验证码" />
          <img src={form.captcha} alt="captcha" onClick={refreshCaptcha} />
        </Form.Item>

        <Form.Item>
          <Button block type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}
