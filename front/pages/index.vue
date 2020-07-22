<template>
  <div class="page-box">
    <h3>最新主题</h3>

    <div class="wiki-list">
      <div v-for="item in wikiList" :key="item.id" class="wiki-item">
        <div class="title">
          <nuxt-link :to="`wiki/${item.number}`">{{ item.title }}</nuxt-link>
        </div>
        <div class="info">
          <img :src="item.user.avatar_url" alt="item.user.login" />
          <span>{{ item.user.login }}</span>
          <span>创建于</span>
          <span>{{ item.created_at }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios, error }) {
    try {
      const wikiList = await $axios.get('/api-v3/repos/justyeh/zrbhsc/issues')
      return { wikiList }
    } catch (e) {
      console.error(e)
      error({ statusCode: 500, message: 'get list error' })
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
.wiki-item {
  .title {
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 10px;
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
}
.wiki-item + .wiki-item {
  margin-top: 30px;
}
</style>
