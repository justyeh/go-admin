import React from 'react'
import AntIcon from '@/components/AntIcon'
import { Link } from 'react-router-dom'
import { Button } from 'antd'
import './style.scss'

const applitions = [
  {
    key: '/application/image',
    icon: 'FileImageOutlined',
    label: '图片库'
  },
  {
    key: '/application/video',
    icon: 'VideoCameraAddOutlined',
    label: '视频库'
  },
  {
    key: '/application/file',
    icon: 'FileZipOutlined',
    label: '文件库'
  },
  {
    key: '/application/form',
    icon: 'FormOutlined',
    label: '万能表单'
  },
  {
    key: '/application/taboo',
    icon: 'DisconnectOutlined',
    label: '敏感词'
  },
  {
    key: '/application/ipblacklist',
    icon: 'EyeInvisibleOutlined',
    label: 'IP黑名单'
  },
  {
    key: '/application/notification',
    icon: 'NotificationOutlined',
    label: '业务提醒'
  },
  {
    key: '/application/operation-log',
    icon: 'SolutionOutlined',
    label: '操作日志'
  }
]
export default function NotFount() {
  return (
    <div className="applition-home-page">
      <h3>应用中心</h3>
      <div>大企业，有大格局！国内首家针对集团与门户需求，以平台+应用方式扩展您的需求，开启您的想象！</div>
      <div className="wrap">
        {applitions.map((item) => (
          <Link key={item.key} to={item.key}>
            <div className="icon">
              <AntIcon name={item.icon} />
            </div>
            <span className="label">{item.label}</span>
            <Button type="primary">进入</Button>
          </Link>
        ))}
      </div>
    </div>
  )
}
