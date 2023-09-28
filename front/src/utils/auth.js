// import { createMiddleware } from 'pinia'
import { defineStore } from 'pinia'

// //全局路由守卫
// export const authMiddleware = createMiddleware(
//     ({ store }) => ({
//         //在路由跳转前执行
//         beforeRouteEnter(to, from, next) {
//             const isAuthenticated = store.auth.isAuthenticated
//             if (isAuthenticated) {
//                 next()
//             } else {
//                 next('/login')
//             }
//         }
//     })
// )

//验证模块
export const useAuthStore = defineStore(
    'auth', {
    state: () => ({
        isAuthenticated: false,
    }),
    actions: {
        setIsAuthenticated(value) {
            this.isAuthenticated = value
            console.log(this.isAuthenticated)
        }
    }
})

