<template>
  <div class="page-box">
    <h3>{{ detail.title }}</h3>
    <div class="info">
      <img :src="detail.user.avatar_url" alt="detail.user.login" />
      <span>{{ detail.user.login }}</span>
      <span>创建于</span>
      <span>{{ detail.created_at }}</span>
    </div>
    {{ detail }}
    <div class="markdown-body" v-html="detail.body"></div>
  </div>
</template>

<script>
import 'github-markdown-css/github-markdown.css'
// https://github.com/sindresorhus/github-markdown-css
// https://developer.github.com/v3/markdown/
export default {
  async asyncData({ $axios, params, error }) {
    try {
      const detail = await $axios.get(
        `/api-v3/repos/justyeh/zrbhsc/issues/${params.id}`
      )
      const contentRes = await $axios.$post('/api-v3/markdown', {
        text: detail.body,
      })
      console.log(contentRes)
      return { detail }
    } catch (e) {
      console.error(e)
      error({ statusCode: 500, message: 'get detail error' })
    }
  },
}
</script>

<style lang="less" scoped>
.page-box {
  padding: 30px;
  > h3 {
    margin: 0 0 30px 0;
  }
}

.info {
  color: #999;
  display: flex;
  align-items: center;
  span + span {
    margin-left: 10px;
  }
  img {
    width: 40px;
    height: 40px;
    border-radius: 3px;
  }
}
</style>
