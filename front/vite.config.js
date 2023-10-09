import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'




// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {

    proxy: {
      "/front": {
        target: "http://localhost:8888",
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/front/, ""),
      },
      '/article': {
        target: 'http://localhost:8888',
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
})
