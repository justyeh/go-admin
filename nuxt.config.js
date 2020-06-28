export default {
  mode: 'universal',
  target: 'server',
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  // 全局css
  css: ['@/assets/style/index.less'],
  plugins: ['~/plugins/axios'],
  components: true,
  buildModules: ['@nuxtjs/eslint-module', '@nuxtjs/stylelint-module'],
  // nuxt module
  modules: ['@nuxtjs/axios'],
  // Axios module configuration
  axios: {
    proxy: true,
  },
  // Build configuration
  build: {},
  // 反向代理
  proxy: {
    '/api-v3/': {
      target: 'https://api.github.com/',
      pathRewrite: { '^/api-v3/': '/' },
    },
  },
  // loading配置
  loading: {
    color: '#04acf7',
    height: '4px',
  },
}
