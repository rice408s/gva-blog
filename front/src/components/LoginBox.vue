<template>
    <div class="login-box">
        <h2>登录</h2>
        <el-form>
            <el-form-item label="用户名">
                <el-input v-model="form.username" placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码">
                <el-input type="password" v-model="form.password" placeholder="请输入密码"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="sub" @click="submitForm">登录</el-button>
            </el-form-item>
        </el-form>
        <div class="register-link">
            还没有账号？<router-link to="/register">注册</router-link>
        </div>
    </div>
</template>

<script  setup>
import { reactive } from 'vue'

import { userUserStore } from '../pinia/modules/user'

// const router = useRouter()

const form = reactive({
    username: '',
    password: ''
})

const userStore = userUserStore()

const submitForm = async () => {
    await userStore.login(form)
}

</script>

<style scoped>
.el-button[type="sub"] {
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.login-box {
    max-width: 400px;
    margin: 0 auto;
    padding: 2rem;
    border: 1px solid #ccc;
    border-radius: 4px;
}

h2 {
    margin-top: 0;
}

.el-input {
    width: 100%;
    /* border: 1px solid #ccc; */
    /* color: #007bff;
     */

}

.register-link {
    margin-top: 1rem;
    text-align: center;
}
</style>