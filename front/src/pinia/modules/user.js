
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

                this.username = userData.username
                const result = await axios.post('/front/login', userData)
                const { data, code } = result.data
                if (code === 0) {
                    this.token = data.token
                    authStore.setIsAuthenticated(true)
                    
                    
                    localStorage.setItem('token', this.token)
                    
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
            localStorage.removeItem('token')
        },
    
        async init(){
            const token = localStorage.getItem('token')
            if (token){
                console.log(token)
                try{
                    const res=await  axios.post('/front/verify',null,{
                        headers:{
                            Authorization:`Bearer ${token}`
                        }
                    })
                    const {code ,data}=res.data
                    if (code==0){
                        this.username=data.username
                        this.token=token
                        this.userId=data.userId
                    }else{
                        localStorage.removeItem('token')
                    }
                }catch(e){
                    console.error(e)
                }
            }
        }
    
    },

    
})





