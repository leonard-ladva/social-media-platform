import { createRouter, createWebHistory } from 'vue-router'
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
		redirect: {name: 'feed'},
		component: HomeView,
		children: [
			{
				path: 'feed',
				name: 'feed',
				component: PostsView,
			},
			{
				path: 'chat/:id',
				name: 'chat', 
				component: ChatsView,
			}
		]
	},
	{
		path: '/login',
		name: 'login',
		component: ()=> import('../components/Login.vue'),
	
	},
	{
		path: '/register',
		name: 'register',
		component: ()=> import('../components/Register.vue'),
	}
]

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
})


export default router
