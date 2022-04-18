import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createStore} from 'vuex'

const store = createStore({
	state () {
		return {
			user: null,
		}
	},
	getters: {
		user: (state) => {
			return state.user
		}
	},
	actions: {
		user(context, user) {
			context.commit('user', user)
		}
	},
	mutations: {
		user(state, user) {
			state.user = user
		}
	},
})


createApp(App)
.use(router)
.use(store)
.mount('#app')

