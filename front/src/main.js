import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
// import 'element-plus/lib/theme-chalk/index.css'
import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import  {userUserStore}  from './pinia/modules/user'

const app = createApp(App)

// Register all icons as components
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(createPinia())
app.use(router)

// //初始化用户名
const userStore = userUserStore()
console.log("hello")

// console.log(userStore)

userStore.init()

app.use(ElementPlus)
app.mount('#app')
