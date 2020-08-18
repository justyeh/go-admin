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
    layout: 'main'
  },
  {
    path: '*',
    title: '404',
    component: 'pages/features/NotFount',
    layout: 'fullpage'
  }
]
