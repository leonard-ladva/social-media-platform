import { createRouter, createWebHistory } from 'vue-router'
import LoginForm from '../components/Login.vue'
import RegisterForm from '../components/Register.vue'
import HomeView from '../components/Home.vue'

const routes = [
	{
		path: '/',
		name: 'home',
		component: HomeView,
	},
	{
		path: '/login',
		name: 'login',
		component: LoginForm,
	},
	{
		path: '/register',
		name: 'register',
		component: RegisterForm,
	}
]

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
})

export default router
