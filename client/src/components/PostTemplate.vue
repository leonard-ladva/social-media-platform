<template>
	<div class="post">
		<div class="header">
			<span class="profilePicture" :style="style"></span>
			<p class="nickname">{{ user.nickname }} </p>
			<p>{{ dateTime }}</p>
		</div>
		<div class="body">
			<p>{{ post.content }}</p>
			<p>{{ tag.title }}</p>
		</div>
	</div>
</template>

<script>
export default {
	props: {
		post: Object,
		user: Object,
		tag: Object,
	},
	computed: {
		style() {
			return 'background-color: ' + this.user.color
		}, 
		dateTime() {
			let date = new Date(this.post.createdAt)
			let now =  new Date()

			let month = date.toLocaleString("default", {month: "short"})
			let day = date.getDate()

			const isSameDay = (first, second) => {
				return first.getFullYear() === second.getFullYear() && 
				first.getMonth() === second.getMonth() && 
				first.getDate() === second.getDate()
			}
			return isSameDay(date, now) ? (Math.floor((now.getTime() - date.getTime()) / 3600000)).toString() + 'h' : month + ' ' + day
		}
	}
}
</script>
