import React, { useState, useCallback } from 'react'
import { useEffectOnce } from 'react-use'
import { Button, Input, Form, Tooltip, Spin } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import './style.scss'

import { getCaptcha, login } from '@/apis/auth'

export default () => {
  const [form, setForm] = useState({ uuid: '', captchaImage: '' })
  const [captchaLoading, setCaptchaLoading] = useState(false)
  const refreshCaptcha = useCallback(async () => {
    setCaptchaLoading(true)
    try {
      const { uuid, image } = await getCaptcha()
      setForm({ uuid, captchaImage: image })
    } catch (error) {}
    setCaptchaLoading(false)
  }, [])

  useEffectOnce(() => {
    refreshCaptcha()
  }, [])

  const handleSubmit = async (values) => {
    try {
      await login({ ...values, uuid: form.uuid })
    } catch (error) {}
  }

  return (
    <div className="login-page">
      <Form size="large" name="login" onFinish={handleSubmit}>
        <h1>G-CMS</h1>

        <Form.Item name="account">
          <Input placeholder="登录账户" prefix={<UserOutlined />} />
        </Form.Item>

        <Form.Item name="password" rules={[{ required: true, message: '请输入密码!' }]}>
          <Input.Password placeholder="登录密码" prefix={<LockOutlined />} />
        </Form.Item>

        <div className="captcha-item">
          <Form.Item name="captcha">
            <Input placeholder="验证码" />
          </Form.Item>
          <Tooltip placement="top" title="刷新验证码">
            <Spin spinning={captchaLoading}>
              <img src={form.captchaImage} alt="captcha" onClick={refreshCaptcha} />
            </Spin>
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
