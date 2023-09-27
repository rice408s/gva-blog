import { ref } from 'vue'
import { defineStore } from 'pinia'


export const userUserStore = defineStore(
'user', () => {
        const userInfo = ref({
        ID: '',
        nickname: '',
        headerImg: '',
        authorty: '',
        // sideMode:''
        activeColor: 'var(--el-color-primary)',
        baseColor: '#fff'
    })
        const token = ref(window.localStore.getItem('token') || '')

    const setUserInfo = (val) => {
        userInfo.value = val
    }


const setToken = (val) => {
    token.value = val
}

const ResetUserInfo = (value = {}) => {
    userInfo.value = {
        // 追加
        ...userInfo.value,
        ...value
    }

}
}
)


