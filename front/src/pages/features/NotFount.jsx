import React from 'react'
import Image404 from '@/assets/image/404.png'
import ImageCloud from '@/assets/image/404_cloud.png'
import { Link } from 'react-router-dom'
import './style.scss'

export default function NotFount() {
  return (
    <div className="not-found-page">
      <div className="wrap">
        <div className="pic">
          <img className="pic__parent" src={Image404} alt="" />
          <img className="pic__child left" src={ImageCloud} alt="" />
          <img className="pic__child mid" src={ImageCloud} alt="" />
          <img className="pic__child right" src={ImageCloud} alt="" />
        </div>
        <div className="bullshit">
          <div className="bullshit__oops">OOPS!</div>
          <div className="bullshit__headline">网管说这个页面你不能进......</div>
          <div className="bullshit__info">请检查您输入的网址是否正确，请点击以下按钮返回主页或者发送错误报告</div>
          <Link to="/" className="bullshit__return-home">
            返回首页
          </Link>
        </div>
      </div> 
    </div>
  )
}
