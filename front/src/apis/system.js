import request from '@/utils/request'

/* 菜单 */
export function menuTree(params) {
  return request.get('/api/system/menu/tree', { params })
}

export function addMenu(data) {
  return request.post('/api/system/menu', data)
}

export function delMenu(id) {
  return request.delete('/api/system/menu/' + id)
}

export function editMenu(data) {
  return request.put('/api/system/menu', data)
}

/* 部门 */

export function deptTree(params) {
  return request.get('/api/system/dept/tree', { params })
}

export function addDept(data) {
  return request.post('/api/system/dept', data)
}

export function delDept(id) {
  return request.delete('/api/system/dept/' + id)
}

export function editDept(data) {
  return request.put('/api/system/dept', data)
}

/* 岗位 */

export function jobTree(params) {
  return request.get('/api/system/job/tree', { params })
}

export function addJob(data) {
  return request.post('/api/system/job', data)
}

export function delJob(id) {
  return request.delete('/api/system/job/' + id)
}

export function editJob(data) {
  return request.put('/api/system/job', data)
}