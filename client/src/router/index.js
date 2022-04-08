import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import HomeSmth from '../components/Home.vue'
import LoginForm from '../components/Login.vue'
import RegisterForm from '../components/Register.vue'

const routes = [
	{
		path: '/',
		name: 'home',
		component: HomeSmth
	},
	{
		path: '/about',
		name: 'about',
		component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
	},
	{
		path: '/feed',
		name: 'feed',
		component: () => import('../views/FeedView.vue')
	},
	{
		path: '/makepost',
		name: 'new post',
		component: () => import('../views/MakePost.vue')
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
