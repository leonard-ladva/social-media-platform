import { createRouter, createWebHistory } from 'vue-router'
import LoginForm from '../components/Login.vue'
import RegisterForm from '../components/Register.vue'
import HomeView from '../components/Home.vue'
import ChatsView from '../components/ChatsView.vue'
import PostsView from '../components/PostsView.vue'

const routes = [
	{
		path: '/:pathMatch(.*)*',
		name: 'NotFound',
		component: ()=> import('../components/NotFound.vue')
	},
	{
		path: '/',
		name: 'home',
		redirect: '/feed',
		component: HomeView,
		children: [
			{
				path: '/feed',
				name: 'feed',
				component: PostsView,
			},
			{
				path: '/chat/:id',
				name: 'chat', 
				component: ChatsView,
			}
		]
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
