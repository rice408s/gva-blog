<template>
  <div class="home">
    <h1>登陆成功主页</h1>
    <el-form-item>
      <el-button type="sub" @click="logout">注销</el-button>
    </el-form-item>
  </div>
</template>



<script setup>
import { RouterView } from 'vue-router'

import { userUserStore } from '../pinia/modules/user'

const userStore = userUserStore()


const logout = async () => {
  await userStore.logout()
}

</script>