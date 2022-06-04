export { chatID }

// chatID returns receiver and sender IDs concatenated in alphabetical order
const chatID = (senderID, receiverID) => {
	return senderID < receiverID ? senderID + receiverID : receiverID + senderID
}