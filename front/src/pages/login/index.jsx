import React, { useState, useCallback } from 'react'
import { useMount } from 'react-use'
import { useHistory } from 'react-router-dom'
import { Button, Input, Form, Tooltip, Spin } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import md5 from 'blueimp-md5'
import { getQueryVariable } from '@/utils/index'
import './style.scss'

import { getCaptcha, login } from '@/apis/auth'

export default () => {
  const history = useHistory()

  const [form, setForm] = useState({ uuid: '', captchaImage: '' })
  const [submitLoading, setsubmitLoading] = useState(false)
  const [captchaLoading, setCaptchaLoading] = useState(false)

  const refreshCaptcha = useCallback(async () => {
    setCaptchaLoading(true)
    try {
      const { uuid, image } = await getCaptcha()
      setForm({ uuid, captchaImage: image })
    } catch (error) {}
    setCaptchaLoading(false)
  }, [])

  useMount(() => {
    refreshCaptcha()
  })

  const handleSubmit = async (values) => {
    setsubmitLoading(true)
    try {
      values.password = md5(values.password)
      const { token } = await login({ ...values, uuid: form.uuid })
      localStorage.setItem('token', token)
      history.replace(getQueryVariable('redirect') || '/')
    } catch (error) {}
    setsubmitLoading(false)
  }

  return (
    <div className="login-page">
      <Form size="large" name="login" onFinish={handleSubmit}>
        <h1>G-CMS</h1>

        <Form.Item name="account" rules={[{ required: true, message: '请输入登陆账户' }]}>
          <Input placeholder="登录账户" prefix={<UserOutlined />} />
        </Form.Item>

        <Form.Item name="password" rules={[{ required: true, message: '请输入登陆密码' }]}>
          <Input.Password placeholder="登录密码" prefix={<LockOutlined />} />
        </Form.Item>

        <div className="captcha-item">
          <Form.Item name="captcha" rules={[{ required: true, message: '请输入验证码' }]}>
            <Input placeholder="验证码" />
          </Form.Item>
          <Tooltip placement="top" title="刷新验证码">
            <Spin spinning={captchaLoading}>
              <img src={form.captchaImage} alt="captcha" onClick={refreshCaptcha} />
            </Spin>
          </Tooltip>
        </div>

        <Form.Item style={{ paddingTop: 10 }}>
          <Button block type="primary" htmlType="submit" loading={submitLoading}>
            登陆
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}
