<template>
	<div class="userCard">
		<span class="profilePicture" :style="profilePic" ></span>
		<span class="nickname">{{ user.nickname }}</span>
	</div>
</template>

<script>
	export default {
		name: "UserCard",
		props: {
			user: Object,
		},
		computed: {
			profilePic() {
				return 'background-color: ' + this.user.color
			}, 
		},
	}
</script>