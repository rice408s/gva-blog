<script setup>
import LoginBox from '../components/LoginBox.vue';
// import { RouterLink } from 'vue-router'
// import HelloWorld from '../components/HelloWorld.vue'
</script>

<template>
    <!-- <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" /> -->

    <!-- <HelloWorld class="hello" msg="这是一个博客" /> -->



    <div class="common-layout">
        <el-container>
            <el-main>
                <LoginBox />
            </el-main>
        </el-container>


    </div>
</template>




<style scoped>
.common-layout {
    overflow-y: auto;
    padding-right: calc(100% - 100vw);

    /* 假设滚动条宽度为默认值 */
    ::-webkit-scrollbar {
        width: 16px;
        /* 假设滚动条宽度为16px */
    }

    /* background-color: #d6e1f3; */
    /* height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    margin-left: 0; */
}</style>