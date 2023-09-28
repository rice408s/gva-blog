
import { defineStore } from 'pinia'
import axios from 'axios'
import { useAuthStore } from '../../utils/auth.js'
import router from '../../router/index'



export const userUserStore = defineStore(
    'user', {
    //状态函数
    state: () => ({
        username: '',
        token: '',
        userId: '',
    }),

    getters: {
        hello: state => 'hello:' + state.username
    },

    //actions
    actions: {
        //异步登录
        async login(userData) {

            try {
                const authStore = useAuthStore()
                console.log(authStore)

                this.username = userData.username
                const result = await axios.post('/front/login', userData)
                const { data, code } = result.data
                if (code === 0) {
                    this.token = data.token
                    console.log(this.username, this.token, code)
                    authStore.setIsAuthenticated(true)
                    console.log(authStore.isAuthenticated)
                    localStorage.setItem('isAuthenticated', true)
                    router.push('/home')
                }
                else {
                    alert('登录失败')
                }
            } catch (e) {
                console.log(e)
            }
        },

        //同步注销
        logout() {
            const authStore = useAuthStore()
            this.username = ''
            this.token = ''
            this.userId = ''
            authStore.setIsAuthenticated(false)
            console.log(authStore.isAuthenticated)
            router.push('/login')
        }
    }
})





