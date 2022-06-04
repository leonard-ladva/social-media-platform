import axios from './axios.js'
import { ws } from './websocket.js'
import { createStore } from 'vuex'
import { useToast } from 'vue-toastification'
import Template_Notification from '../components/Template_Notification.vue'

export { store }

const toast = useToast()

const store = createStore({
	state () {
		return {
			currentUser: null,
			loggedIn: false,
			allUsers: new Map(), // UserID: User
			newMessages: new Map(), // ChatID: message
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
			return new Promise((resolve, reject) => {
				axios.get('/user')
				.then (({data, status}) => {
					if (status === 200) {
						context.commit('setLoggedIn')
						context.commit('setCurrentUser', data)
						ws.connect(data)

						resolve(true)
					}
				})
				.catch((error) => {
					console.log('got error')
					reject(error)
				}) 
			})
		},
		async logInUser(context, user) {
			return new Promise((resolve, reject) => {
				axios.post('/login', {
					nickname: user.nickname,
					passwordPlain: user.password,
				})
				.then(({data, status}) => {
					if (status === 200) {
						localStorage.setItem('token', data.token)
						context.commit('setLoggedIn')
						context.commit('setCurrentUser', data.user)
						resolve(true)
					}
				})
				.catch (error => {
					reject(error)
				})
			})
		},
		logOutUser(context) {
			localStorage.removeItem('token')
			context.commit('setLoggedOut')
			context.commit('setCurrentUser', null)
			ws.disconnect()
		},
		newMessage({commit, state}, message) {
			commit('addMessage', message)
			if (message.userId === state.currentUser.id) { return }
			const sender = state.allUsers.get(message.userId)
			toast({
				component: Template_Notification,
				props: {
					message: message,
					sender: sender,
				}
			})
		},
		removeFirstNotification(context) {
			context.commit('removeFirstNotification')
		},
		userWentOffline(context, userID) {
			context.commit('setUserNotActive', userID)
		},
		userCameOnline(context, userID) {
			context.commit('setUserActive', userID)
		}
	},
	mutations: {
		newUser(state, user) {
			state.allUsers.set(user.id, user)
		},
		setUserNotActive(state, userID) {
			state.allUsers.get(userID).active = false
		},
		setUserActive(state, userID) {
			state.allUsers.get(userID).active = true
		},
		addMessage(state, message) {
			state.newMessages.set(message.chatId, message)
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
	},
})