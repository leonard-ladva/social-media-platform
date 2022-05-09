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
		}
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
		}
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
	},
})


createApp(App)
.use(router)
.use(store)
.mount('#app')

