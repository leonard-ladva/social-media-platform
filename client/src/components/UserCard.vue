<template>
	<router-link :to="{name: 'chat', params:{id: user.id}}" class="chatLink"> 
		<div class="userCard">
			<span class="profilePicture" :style="profilePic" ></span>
			<span class="nickname">{{ user.nickname }}</span>
		</div>
	</router-link>
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
</script>:
<style>
	.userCard {
		background-color: rgb(239, 243, 244);
		height: 55px;
		border-radius: 30px;
		display: flex;
		align-items: center;
	}
	.userCard .profilePicture {
		height: 45px;
		width: 45px;	
		margin: 5px;
	}
	.userCard .nickname {
		width: calc(100% - 80px);
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		font-size: 20px;
		padding-bottom: 3px;
		color: black;
		text-decoration: none;
	}
	.userCard:hover {
		background-color: #E1E8ED; 
	}
	.chatLink {
		text-decoration: none;
	}

</style>