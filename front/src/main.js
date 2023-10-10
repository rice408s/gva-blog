import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import Antd from 'ant-design-vue';

import 'ant-design-vue/dist/reset.css';
import  {userUserStore}  from './pinia/modules/user'

const app = createApp(App)


for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(createPinia())
app.use(router)

const userStore = userUserStore()
userStore.init()
app.use(Antd)
app.use(ElementPlus)
app.mount('#app')
