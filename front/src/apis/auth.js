import request from '@/utils/request'

export function getCaptcha() {
  return request.get('/api/auth/captcha?w=160&h=40')
}

export function login(data) {
  return request.post('/api/auth/login', data)
}
