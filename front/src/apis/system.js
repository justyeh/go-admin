import request from '@/utils/request'

export function menuTree(params) {
  return request.get('/api/system/menu/tree', { params })
}

export function addMenu(data) {
  return request.post('/api/system/menu', data)
}

export function delMenu(id) {
  return request.delete('/api/system/menu?id=' + id)
}

export function updateMenu(data) {
  return request.put('/api/system/menu', data)
}
