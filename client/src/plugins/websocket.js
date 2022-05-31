export { ws }

import { store } from '../main.js'

const wsAddress = "ws://localhost:9100/ws"

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
					console.log('Received')
					store.dispatch('newMessage', true)
					if (wsMsg.message.userId != store.state.user.id) {
						store.dispatch('addMessage', wsMsg.message)	
					}
					break
				// A new user came online
				case 'online':
				
					break
				// A user went offline
				case 'offline':

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
	sendMessage: function(message) {
		let wsMsg = {type: 'message', message: message}
		let msg = new WsMsg(wsMsg)
		this.connection.send(JSON.stringify(msg))	
	}
}