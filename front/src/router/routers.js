export default [
  {
    path: '/',
    title: '首页',
    exact: true,
    component: 'pages/home/index',
    layout: 'management'
  },
  {
    path: '/login',
    title: '登录',
    component: 'pages/login/index',
    layout: 'fullpage'
  },
  {
    path: '/profile',
    exact: true,
    title: '个人中心',
    component: 'pages/profile/index',
    layout: 'management'
  },

  {
    path: '/system/user',
    exact: true,
    title: '用户管理',
    component: 'pages/system/user/index',
    layout: 'management'
  },
  {
    path: '/system/role',
    exact: true,
    title: '角色管理',
    component: 'pages/system/role/index',
    layout: 'management'
  },
  {
    path: '/system/permission',
    exact: true,
    title: '权限管理',
    component: 'pages/system/permission/index',
    layout: 'management'
  },
  {
    path: '/system/menu',
    exact: true,
    title: '菜单管理',
    component: 'pages/system/menu/index',
    layout: 'management'
  },
  {
    path: '/system/dept',
    exact: true,
    title: '部门管理',
    component: 'pages/system/dept/index',
    layout: 'management'
  },
  {
    path: '/system/job',
    exact: true,
    title: '岗位管理',
    component: 'pages/system/job/index',
    layout: 'management'
  },
  {
    path: '/system/dictionary',
    exact: true,
    title: '字典管理',
    component: 'pages/system/dictionary/index',
    layout: 'management'
  },

  {
    path: '*',
    title: '404',
    component: 'pages/features/NotFount',
    layout: 'fullpage'
  }
]
