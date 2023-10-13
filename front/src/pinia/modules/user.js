
import { defineStore } from 'pinia'
import axios from 'axios'
import { useAuthStore } from '../../utils/auth.js'
import router from '../../router/index'



export const userUserStore = defineStore(
    'user', {
    //状态函数
    state: () => ({
        username: '',
        nickName: '',
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


                const result = await axios.post('/front/login', userData)
                const { data, code } = result.data
                if (code === 0) {
                    console.log(data)
                    this.token = data.token
                    authStore.setIsAuthenticated(true)
                    this.username = userData.username
                    this.nickName = data.user.nickName
                    console.log(this.nickName)
                    localStorage.setItem('token', this.token)
                    
                    router.push('/home')
                }
                else {
                    alert('登录失败')
                    this.username=''
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
                console.log('good')
                console.log(token)
                try{
                    const res=await  axios.post('/front/verify',null,{
                        headers:{
                            Authorization:`${token}`
                        }
                    })
                    const {code ,user}=res.data
                    console.log('xxxxxx')
                    console.log(res.data)
                    console.log(user.ID)
                    if (code==0){
                        this.username=user.userName
                        this.token=token
                        this.userId=user.ID
                    }else{
                        localStorage.removeItem('token')
                        console.log('bad')
                    }
                }catch(e){
                    console.error(e)
                }
            }
        }
    
    },

    
})





