import { createRouter, createWebHashHistory } from 'vue-router'
import { userUserStore } from '@/pinia/modules/user'
import { store } from '@/pinia/index'
import { useAuthStore } from '@/utils/auth'
// import { useAuthStore } from '../utils/auth'


const router = createRouter({
  history: createWebHashHistory(/*import.meta.env.BASE_URL*/),
  routes: [
    {
      path: '/',
      redirect: "/home"
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
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginPage.vue')
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
    }, {
      path: '/article/:id',
      name: 'ArticleDetail',
      component: () => import('../views/ArticleDetailView.vue'),
      props: true
    },
  ]
})



router.beforeEach(async (to, from, next) => {
  // const store = userUserStore()
  const store = useAuthStore()
  console.log(store.isAuthenticated)
  if (to.meta.requiresAuth) {
    if (store.isAuthenticated) {
      next()
    } else {
      next('/login')
    }
  }
  next()
})

export default router
