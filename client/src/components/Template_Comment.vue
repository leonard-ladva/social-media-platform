<template>
	<div class="comment">
		<div class="header">
				<span class="profilePicture" :style="style"></span>
				<h6 class="user">{{ user.firstName }} {{ user.lastName }} 
					<span>@{{user.nickname}}</span>
					<span>· {{dateTime}}</span>	
				</h6>
			</div>
			<div class="body">
				<p class="content">{{ comment.content }}</p>
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
		padding: 1.2rem;
	}
	.comment .header {
		display: flex;
	}
	.comment .body {
		text-align: left;
		margin: -1rem 0 0 calc(50px + 1rem);
	}
	.user {
		margin: 0.2rem 1rem;
	}
	.user span {
		color: var(--darkGrey);
		font-family: "Chirp Regular";
		margin-left: 0.4rem;
	}
</style>