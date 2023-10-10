import { useAuthStore } from "@/utils/auth";
import axios from "axios";

const authStore =useAuthStore()

const token = localStorage.getItem('token')

if (token) {
    authStore.setIsAuthenticated(true)
}


axios.interceptors.request.use(
    config => {
        if (token){
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    error => {
        return Promise.reject(error)
    }
)
