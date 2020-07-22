<template>
  <div>
    <div class="form">
      <div class="form-item">
        <span>Account</span>
        <input v-model="account" type="text" />
      </div>
      <div class="form-item">
        <span>Password</span>
        <input v-model="password" type="password" />
      </div>
      <div class="form-item">
        <span></span>
        <button @click="login">Login</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  layout: 'fullpage',
  data() {
    return {
      account: '',
      password: '',
    }
  },
  methods: {
    login() {},
  },
}
</script>
