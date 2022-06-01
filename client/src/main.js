import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createStore } from 'vuex'
import './assets/css/style.css'
import axios from 'axios'

export { store }

const store = createStore({
	state () {
		return {
			currentUser: null,
			loggedIn: false,
			allUsers: new Map(), // UserID: UserObject
			newMessages: new Map(), // ChatID: message
			notifications: []
		}
	},
	getters: {
		activeUsers(state) {
			let activeUsers = new Map()
			for (let user of state.allUsers.values()) {
				if (user.active === true) {
					activeUsers.set(user.id, user)
				}
			}
			return activeUsers
		},
		offlineUsers(state) {
			let offlineUsers = new Map()
			for (let user of state.allUsers.values()) {
				if (user.active === false) {
					offlineUsers.set(user.id, user)
				}
			}
			return offlineUsers
		},
		loaded(state) {
			return state.currentUser && state.allUsers.size != 0
		}
	},
	actions: {
		async getUsers(context) {
			let response = await axios.get('/users')	
			if (response.status === 200) {
				for (let user of response.data) {
					context.commit('newUser', user)
				}
			} else {
				console.log(`ERROR: getting users. Status: ${response.status}`)
			}
		},
		async getCurrentUser(context) {
			let response = await axios.get('/user')
			if (response.status === 200) {
				context.commit('setLoggedIn')
				context.commit('setCurrentUser', response.data)
			} else if (response.status === 401) {
				console.log("You're Not Logged In.")
			} else {
				console.log(`ERROR: getting Current User. Status: ${response.status}`)
			}
		},
		logInUser(context, user) {
			context.commit('setLoggedIn')
			context.commit('setCurrentUser', user)
		},
		logOutUser(context) {
			context.commit('setLoggedOut')
			context.commit('setCurrentUser', null)
		},
		newMessage(context, message) {
			context.commit('addMessage', message)
			context.commit('newNotification', message)
		},
		removeFirstNotification(context) {
			context.commit('removeFirstNotification')
		}
	},
	mutations: {
		newUser(state, user) {
			state.allUsers.set(user.id, user)
		},
		addMessage(state, message) {
			state.newMessages.set(message.chatId, message)
		},
		newMessage(state, haveNew) {
			state.newMessage = haveNew
		},
		setLoggedIn(state) {
			state.loggedIn = true
		},
		setLoggedOut(state) {
			state.loggedIn = false
		},
		setCurrentUser(state, user) {
			state.currentUser = user
		},
		newNotification(state, notification) {
			state.notifications.push(notification)
		},
		removeFirstNotification(state) {
			state.notifications.shift()
		}
	},
})


createApp(App)
.use(router)
.use(store)
.mount('#app')
