import { createRouter, createWebHistory } from 'vue-router'
import LoginForm from '../components/Login.vue'
import RegisterForm from '../components/Register.vue'

const routes = [
	{
		path: '/',
		name: 'home',
		component: () => import('../views/FeedView.vue')
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
