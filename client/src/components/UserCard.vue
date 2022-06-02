<template>
	<router-link :to="{name: 'chat', params:{receiverId: user.id}}" class="chatLink"> 
		<div class="userCard" :style="background">
			<span class="profilePicture" :style="profilePic" ></span>
			<span class="nickname">{{ user.nickname }}</span>
			<span v-if="active" class="activeDot"></span>
		</div>
	</router-link>
</template>

<script>
	export default {
		name: "UserCard",
		props: {
			user: Object,
			active: Boolean,
		},
		computed: {
			profilePic() {
				return 'background-color: ' + this.user.color
			},
			background() {
				// If current route is this users chat
				if (this.$route.params.receiverId == this.user.id) {
					return 'background-color: var(--extraLightGrey)'
				}
				return 'background-color: var(--extraExtraLightGrey)'
			}
		},
	}
</script>:
<style>
	.userCard {
		background-color: var(--extraExtraLightGrey);
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
		width: calc(100% - 90px);
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		font-size: 20px;
		padding-bottom: 3px;
		color: var(--black);
		text-decoration: none;
	}
	.userCard:hover {
		background-color: var(--extraLightGrey); 
	}
	.chatLink {
		text-decoration: none;
	}
	.activeDot {
		background-color: #5ad032;
		border-radius: 50%;
		height: 13px;
		width: 13px;
	}
</style>