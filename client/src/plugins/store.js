import axios from './axios.js'
import { ws } from './websocket.js'
import { createStore } from 'vuex'
import { useToast } from 'vue-toastification'
import Template_Notification from '../components/Template_Notification.vue'
// import { chatID } from '../assets/js/chats.js'

const toast = useToast()

const store = createStore({
	state () {
		return {
			currentUser: null,
			loggedIn: false,
			allUsers: new Map(), // UserID: User
			newMessages: new Map(), // ChatID: message
			chats: new Map(), // ChatID: chat
		}
	},
	getters: {
		// activeChats(state) {
			// let activeChats = Array.from(state.chats)
		// },
		activeUsers(state) {
			let activeUsers = Array.from(state.allUsers.values())
			// filter out offline users and the current User
			activeUsers = activeUsers.filter(user => user.active === true && user.id !== state.currentUser.id)
			let existingChats = activeUsers.filter(user => user.lastMessageTime !== -1)
			let newChats = activeUsers.filter(user => user.lastMessageTime === -1)

			existingChats = existingChats.sort((a, b) => {
				return b.lastMessageTime - a.lastMessageTime
			})

			newChats = newChats.sort((a, b) => {
				return a.nickname - b.nickname
			})
			activeUsers = existingChats.concat(newChats)

			return activeUsers
		},
		offlineUsers(state) {
			let offlineUsers = Array.from(state.allUsers.values())
			// filter out active users and the current User
			offlineUsers = offlineUsers.filter(user => user.active === false && user.id !== state.currentUser.id)
			let existingChats = offlineUsers.filter(user => user.lastMessageTime !== -1)
			let newChats = offlineUsers.filter(user => user.lastMessageTime === -1)

			existingChats = existingChats.sort((a, b) => {
				return b.lastMessageTime - a.lastMessageTime
			})

			newChats = newChats.sort((a, b) => {
				return a.nickname - b.nickname
			})
			offlineUsers = existingChats.concat(newChats)

			return offlineUsers
		},
		loaded(state) {
			return state.currentUser && state.allUsers.size != 0
		}
	},
	actions: {
		async getChats({commit, state}) {
			axios.get('/chats')	
			.then(response => {
				for (let chat of response.data) {
					let userA = chat.id.substr(0, 36)
					let userB = chat.id.substr(36)
					// Ignore chat if not related to the current user 
					if (state.currentUser.id !== userA && state.currentUser.id !== userB) { continue }

					let otherUserID = state.currentUser.id === userA ? userB : userA
					commit('changeLastMessageTime', {userId: otherUserID, lastMessageTime: chat.lastMessageTime})	
					commit('addChat', chat)
				}
			}).catch(error => {
				console.log(error)
			})
		},
		async getUsers(context) {
			axios.get('/users')	
			.then(resp => {
				for (let user of resp.data) {
					user.lastMessageTime = -1
					context.commit('newUser', user)
				}
			})
			.then(() => {
				context.dispatch('getChats')
			})
			.catch(error => {
				console.log(error)
			})
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
			commit('changeLastMessageTime', {userId: message.receiverId, lastMessageTime: message.createdAt})
			if (message.userId === state.currentUser.id) { return }
			const sender = state.allUsers.get(message.userId)
			toast({
				component: Template_Notification,
				props: {
					message: message,
					sender: sender,
				}
			})
			commit('changeLastMessageTime', {userId: message.userId, lastMessageTime: message.createdAt})
		},
		removeFirstNotification(context) {
			context.commit('removeFirstNotification')
		},
		userWentOffline(context, userID) {
			context.commit('setUserNotActive', userID)
		},
		userCameOnline(context, userID) {
			if (!context.state.allUsers.has(userID)) {
				context.dispatch('getUsers')
				return
			}
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
		addChat(state, chat) {
			state.chats.set(chat.id, chat)
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
		changeLastMessageTime(state, data) {
			state.allUsers.get(data.userId).lastMessageTime = data.lastMessageTime
		}
	},
})

export { store }