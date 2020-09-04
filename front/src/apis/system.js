import request from '@/utils/request'

/* 用户 */
export function userList(params) {
  return request.get('/api/system/user/list', { params })
}

export function addUser(data) {
  return request.post('/api/system/user', data)
}

export function delUser(id) {
  return request.delete('/api/system/user/' + id)
}

export function editUser(data) {
  return request.put('/api/system/user', data)
}

export function updateUserStatus(data) {
  return request.put('/api/system/user/updateUserStatus', data)
}

/* 权限 */
export function roleList(params) {
  return request.get('/api/system/role/list', { params })
}

export function addRole(data) {
  return request.post('/api/system/role', data)
}

export function delRole(id) {
  return request.delete('/api/system/role/' + id)
}

export function editRole(data) {
  return request.put('/api/system/role', data)
}

export function updateRoleStatus(data) {
  return request.put('/api/system/role/updateRoleStatus', data)
}

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

/* 权限 */
export function permissionTree(params) {
  return request.get('/api/system/permission/tree', { params })
}

export function addPermission(data) {
  return request.post('/api/system/permission', data)
}

export function delPermission(id) {
  return request.delete('/api/system/permission/' + id)
}

export function editPermission(data) {
  return request.put('/api/system/permission', data)
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

export function jobList(params) {
  return request.get('/api/system/job/list', { params })
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

/* 字典 */

export function dictionaryList(params) {
  return request.get('/api/system/dictionary/main/list', { params })
}

export function addDictionary(data) {
  return request.post('/api/system/dictionary/main', data)
}

export function delDictionary(id) {
  return request.delete('/api/system/dictionary/main/' + id)
}

export function editDictionary(data) {
  return request.put('/api/system/dictionary/main', data)
}

export function dictionaryDetailList(params) {
  return request.get('/api/system/dictionary/detail/list', { params })
}

export function addDictionaryDetail(data) {
  return request.post('/api/system/dictionary/detail', data)
}

export function delDictionaryDetail(id) {
  return request.delete('/api/system/dictionary/detail/' + id)
}

export function editDictionaryDetail(data) {
  return request.put('/api/system/dictionary/detail', data)
}
