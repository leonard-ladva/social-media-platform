export { ws }

const wsAddress = "ws://localhost:9100/ws"

function WsMsg(params) {
	this.type = params.type
	this.user = params.user
}

const ws = {
	connection: null,
	connect: function(user) {
		console.log("Starting connection to WebSocket Server")
		this.connection = new WebSocket(wsAddress)

		this.connection.onmessage = function(event) {
			console.log(event);
		}

		this.connection.onopen = function(event) {
			// Authenticate the websocket with the backend, send the current logged in user
			let wsMsgParams = {type: 'auth', user: user}
			let authMsg = new WsMsg(wsMsgParams)
			this.send(JSON.stringify(authMsg))

			console.log(event)
			console.log("Successfully connected to the websocket")
		}
	},
	sendMessage: function(message) {
		this.connection.send(message)	
	}
}