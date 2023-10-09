<template>
    <div class="login-box">
        <h2>登录</h2>
        <el-form>
            <div class="form-item">

                <div class="form-input">
                    <el-input v-model.trim="form.username" autocomplete="off" style="width: 100%" placeholder="请输入账号" />
                </div>
            </div>
            <div class="form-item">

                <div class="form-input">
                    <el-input type="password" v-model.trim="form.password" autocomplete="off" style="width: 100%" placeholder="请输入密码" />
                </div>
            </div>
            <el-form-item>
                <el-button type="sub" @click="submitForm" class="login-button">登录</el-button>
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
    console.log(userStore)
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


.form-item {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}

.form-label {
    flex: 0 0 80px;
    text-align: right;
    margin-right: 20px;
}

.form-input {
    flex: 1;
}


h2 {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
    text-align: center;
}

.login-button {
    width: 100%;
    margin-top: 20px;
}
</style>