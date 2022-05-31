import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createStore } from 'vuex'
import './assets/css/style.css'

export { store }

const store = createStore({
	state () {
		return {
			user: null,
			allUsers: null,
			activeUsers: null,
			offlineUsers: null,
			messages: new Map(),
			newMessage: false,
		}
	},
	getters: {
		user: (state) => {
			return state.user
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
		messages: (state) => {
			return state.messages
		}
	},
	actions: {
		user(context, user) {
			context.commit('user', user)
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
		addMessage(context, message) {
			console.log("adding message to store")
			context.commit('addMessage', message)
		},
		newMessage(context, haveNew) {
			console.log('received new message')
			context.commit('newMessage', haveNew)
		}
	},
	mutations: {
		user(state, user) {
			state.user = user
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
		addMessage(state, message) {
			state.messages.set(message.chatId, message)
		},
		newMessage(state, haveNew) {
			state.newMessage = haveNew
		}
	},
})


createApp(App)
.use(router)
.use(store)
.mount('#app')
