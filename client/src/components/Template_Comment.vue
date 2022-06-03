<template>
	<div class="comment">
		<div class="header">
			<span class="profilePicture" :style="style"></span>
			<p class="nickname">{{ user.nickname }} </p>
			<p>{{ dateTime }}</p>
		</div>
		<div class="body">
			<p>{{ comment.content }}</p>
		</div>
	</div>
</template>

<script>
export default {
	props: {
		comment: Object,
		user: Object,
	},
	computed: {
		style() {
			return 'background-color: ' + this.user.color
		}, 
		dateTime() {
			const millisecondsInMinute = 36000
			const millisecondsInHour = 3600000
			let date = new Date(this.comment.createdAt)
			let now =  new Date()

			let month = date.toLocaleString("default", {month: "short"})
			let day = date.getDate()

			const isSameDay = (first, second) => {
				return first.getFullYear() === second.getFullYear() && 
				first.getMonth() === second.getMonth() && 
				first.getDate() === second.getDate()
			}
			const isSameHour = (first, second) => {
				return isSameDay && first.getHours() === second.getHours()
			}
			
			let time = 	!isSameDay(date, now) ?  month + ' ' + day :
						!isSameHour(date, now) ? (Math.floor((now.getTime() - date.getTime()) / millisecondsInHour)).toString() + 'h' : 
						(Math.floor((now.getTime() - date.getTime()) / millisecondsInMinute)).toString() + 'm'
			return time
		}
	}
}
</script>
<style>
	.comment {
		margin: -1px 0 0 -1px ;
		border: 1px solid var(--extraLightGrey);
	}
	.comment .nickname {
		font-family: "Chirp Bold";
	}
	.comment:hover {
		background-color: var(--extraExtraLightGrey);
	}
</style>