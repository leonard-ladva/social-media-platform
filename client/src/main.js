import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createStore } from 'vuex'

const store = createStore({
	state () {
		return {
			user: {},
			tags: null,
		}
	},
	getters: {
		user: (state) => {
			return state.user
		},
		tags: (state) => {
			return state.tags
		}
	},
	actions: {
		user(context, user) {
			context.commit('user', user)
		}, 
		tags(context, tags) {
			context.commit('tags', tags)
		}
	},
	mutations: {
		user(state, user) {
			state.user = user
		},
		tags(state, tags) {
			state.tags = tags
		}
	},
})


createApp(App)
.use(router)
.use(store)
.mount('#app')

