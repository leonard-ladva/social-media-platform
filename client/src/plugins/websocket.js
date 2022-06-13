export { ws }

import { store } from './store.js'

const wsAddress = "ws://localhost:9000/ws"

class WsMsg {
	constructor(params) {
		this.type = params.type
		this.message = params.message
		this.userId = params.userId
	}
}


const ws = {
	connection: null,
	connect: function(user) {
		console.log("Starting connection to WebSocket Server")
		this.connection = new WebSocket(wsAddress)

		this.connection.onmessage = function(event) {
			let wsMsg = JSON.parse(event.data)
			switch (wsMsg.type) {
				// A new message has been sent to the user
				case 'message':
					store.dispatch('newMessage', wsMsg.message)
					break
				// A new user came online
				case 'online':
					if (wsMsg.userId === store.state.currentUser.id) { break }
					store.dispatch('userCameOnline', wsMsg.userId)	
					break
				// A user went offline
				case 'offline':
					store.dispatch('userWentOffline', wsMsg.userId)
					break
			}
		}

		this.connection.onopen = function() {
			// Authenticate the websocket with the backend, send the current logged in user
			let wsMsg = {type: 'auth', userId: user.id}
			let authMsg = new WsMsg(wsMsg)
			this.send(JSON.stringify(authMsg))

			console.log("Successfully connected to the websocket")
		}
	},
	disconnect: function() {
		this.connection.close(1000, "User Loggged out.")
	},
	sendMessage: function(message) {
		let wsMsg = {type: 'message', message: message}
		let msg = new WsMsg(wsMsg)
		this.connection.send(JSON.stringify(msg))	
	}
}