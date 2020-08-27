import React from 'react'
import Image404 from '@/assets/image/404.png'
import { useHistory } from 'react-router-dom'
import { Button } from 'antd'
import './style.scss'

export default function NotFount() {
  const history = useHistory()
  return (
    <div className="not-found-page">
      <div className="wrap">
        <img className="image" src={Image404} alt="" />
        <div className="text">
          <div className="oops">OOPS!</div>
          <div className="headline">网管说这个页面你不能进......</div>
          <div className="info">请检查您输入的网址是否正确，请点击以下按钮返回主页或者发送错误报告</div>
          <Button
            type="primary"
            shape="round"
            onClick={() => {
              history.goBack()
            }}
          >
            返回上级
          </Button>
        </div>
      </div>
    </div>
  )
}
