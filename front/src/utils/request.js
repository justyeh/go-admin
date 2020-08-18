import axios from 'axios'
import { notification } from 'antd'

// 创建axios实例
const service = axios.create({
  baseURL: '/',
  headers: { 'Content-Type': 'application/json' }
})

// request拦截器
service.interceptors.request.use(
  (config) => {
    config.headers['Authorization'] = localStorage.getItem('token')
    return config
  },
  (error) => {
    Promise.reject(error)
  }
)

// response 拦截器
service.interceptors.response.use(
  (response) => {
    const { status, message } = response.data
    if (status === 200) {
      return response.data
    } else {
      notification.error({ message: '请求错误', description: message })
      return Promise.reject(new Error(message))
    }
  },
  (error) => {
    if (error.toString().indexOf('Error: timeout') !== -1) {
      notification.error({ message: 'Error', description: '网络请求超时' })
      return Promise.reject(error)
    }

    notification.error({ message: 'Error', description: '网络请求错误' })
    return Promise.reject(error)
  }
)
export default service
