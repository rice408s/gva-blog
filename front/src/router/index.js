  import { createRouter, createWebHistory } from 'vue-router'

import { useAuthStore } from '../utils/auth'


const router = createRouter({
  history: createWebHistory(/*import.meta.env.BASE_URL*/),
  routes: [
    {
      path:'/',
      redirect :"/login"
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path:'/login',
      name:'login',
      component: () => import('../views/LoginPage.vue')
    },
    {
      path:'/home',
      name:'home',
      component: () => import('../views/HomeView.vue'),
      //需要登录才能访问
      meta:{requiresAuth:true}
    }
  ]
})



// router.beforeEach(())=>{}



// router.beforeEach((to, from, next) => {
//   if (to.meta.requiresAuth) {
//     // 根据你存储登录状态的方式进行验证，比如检查是否存在登录状态的Cookie

//     if (window.localStore.getItem('token')) {
//       next(); // 已登录，继续导航到目标路由
//     } else {
//       next('/login'); // 未登录，重定向到登录页面
//     }
//   } else {
//     next(); // 非需要登录验证的路由，直接导航到目标路由
//   }
// });

export default router
