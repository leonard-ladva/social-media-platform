import { createRouter, createWebHistory } from 'vue-router'

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
		component: ()=> import('../components/Home.vue'),
		children: [
			{
				path: 'feed',
				name: 'feed',
				component: ()=> import('../components/Page_Feed.vue'),
			},
			{
				path: 'chat/:receiverId',
				name: 'chat', 
				component: ()=> import('../components/Page_Chat.vue'),
			},
			{
				path: 'post/:postId',
				name: 'post',
				component: ()=> import('../components/Page_Post.vue'),
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
