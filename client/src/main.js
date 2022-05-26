import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createStore } from 'vuex'
import './assets/css/style.css'

const store = createStore({
	state () {
		return {
			user: null,
			tags: null,
			allUsers: null,
			activeUsers: null,
			offlineUsers: null,
		}
	},
	getters: {
		user: (state) => {
			return state.user
		},
		tags: (state) => {
			return state.tags
		},
		allUsers: (state) => {
			return state.allUsers
		},
		activeUsers: (state) => {
			return state.activeUsers
		},
		offlineUsers: (state) => {
			return state.offlineUsers
		},
	},
	actions: {
		user(context, user) {
			context.commit('user', user)
		}, 
		tags(context, tags) {
			context.commit('tags', tags)
		},
		allUsers(context, allUsers) {
			context.commit('allUsers', allUsers)
		},
		activeUsers(context, activeUsers) {
			context.commit('activeUsers', activeUsers)
		},
		offlineUsers(context, offlineUsers) {
			context.commit('offlineUsers', offlineUsers)
		},
	},
	mutations: {
		user(state, user) {
			state.user = user
		},
		tags(state, tags) {
			state.tags = tags
		},
		allUsers(state, allUsers) {
			state.allUsers = allUsers
		},
		activeUsers(state, activeUsers) {
			state.activeUsers = activeUsers
		},
		offlineUsers(state, offlineUsers) {
			state.offlineUsers = offlineUsers
		},
	},
})


createApp(App)
.use(router)
.use(store)
.mount('#app')
