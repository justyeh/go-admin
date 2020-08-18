import request from '@/utils/request'

export function getCaptcha() {
  return request.get('/api/system/captcha')
}
