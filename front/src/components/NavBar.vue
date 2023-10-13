<template>
    <div class="nav-bar">
        <div class="nav-bar-title">{{ title }}</div>
        <div class="nav-bar-links">
            <div class="search-icon" @click="showSearchBox = showSearchBox == true ? false : true">
                <el-icon>
                    <Search />
                </el-icon>
            </div>
            <!-- todo 搜索框 -->
            <div class="search-box" v-show="showSearchBox">
                <el-input v-model="keyword" placeholder="搜索" @keyup.enter="search" />
            </div>

            <router-link to="/">首页</router-link>
            <router-link to="/archives">归档</router-link>
            <router-link to="/categories">分类</router-link>
            <router-link to="/tags">标签</router-link>
            <router-link to="/links">友链</router-link>
            <router-link to="/about">关于</router-link>
            <el-dropdown v-if="username" trigger="click">
                <span class="el-dropdown-link">
                    {{ nickName }}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>个人信息</el-dropdown-item>
                        <el-dropdown-item>修改密码</el-dropdown-item>
                        <el-dropdown-item divided @click="handleLogout">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
            <router-link v-else to="/login">登录</router-link>
        </div>
    </div>
</template>

<script>
import { defineComponent, ref, computed } from 'vue'
import { RouterLink } from 'vue-router'
import { ElInput } from 'element-plus'
import { userUserStore } from '@/pinia/modules/user'
export default defineComponent({
    name: 'NavBar',
    components: {
        RouterLink,
        ElInput,

    },
    props: {
        title: {
            type: String,
            required: true
        }
    },




    // eslint-disable-next-line no-unused-vars
    setup(props) {
        const keyword = ref('')
        const showSearchBox = ref(false)
        const userStore = userUserStore()

        function search() {
            if (keyword.value) {
                window.location.href = `/search?q=${keyword.value}`
            }
        };

        const handleLogout = () => {
            userStore.logout();
        };
        const nickName = computed(() => userStore.nickName)
        const username = computed(() => userStore.username)
        return {
            keyword,
            showSearchBox,
            search,
            username,
            handleLogout,
            nickName 
        }
    }
})
</script>

<style scoped>
.nav-bar-links {
    flex: 1;
    display: flex;
    justify-content: right;
    align-items: center;
}

.nav-bar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 999;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    background-color: #f5f5f5;
}

.nav-bar-title {
    font-size: 24px;
    font-weight: bold;
}

.nav-bar-links {
    display: flex;
    align-items: right;
}

.nav-bar-links>* {
    margin-left: 30px;
}

.nav-bar-links>*:first-child {
    margin-left: 0;
}

.search-box {
    position: absolute;
    top: 100%;
    right: 0;
    display: none;
    z-index: 1000;
    width: 100%;
    height: 40px;
}

.el-input {
    width: 200px;
}
</style>