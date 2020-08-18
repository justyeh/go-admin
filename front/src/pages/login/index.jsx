import React, { useState, useCallback } from 'react'
import { useEffectOnce } from 'react-use'
import { Button, Input, Form, Tooltip } from 'antd'
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

  useEffectOnce(() => {
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

        <div className="captcha-item">
          <Form.Item name="captcha" rules={[{ required: true, message: '请输入验证码!' }]}>
            <Input placeholder="验证码" />
          </Form.Item>
          <Tooltip placement="top" title="刷新验证码">
            <img src={form.captcha} alt="captcha" onClick={refreshCaptcha} />
          </Tooltip>
        </div>

        <Form.Item>
          <Button block type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}
