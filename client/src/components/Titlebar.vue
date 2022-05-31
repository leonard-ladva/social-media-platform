<template>
	<div id="titlebar">
		<div v-if="$route.name != 'feed'">
			<router-link :to="{name: 'home'}">
				<i id="homeLink" class="iconoir-arrow-left"></i>
			</router-link>
		</div>
		<div>
			<h4>
				{{currentPathName}}
			</h4>
		</div>
	</div>
</template>

<script>
export default {
	name: 'TitleBar',
	computed: {
		currentPathName() {
			return this.$route.name == "feed" ? "Home" : this.$route.name.replace(/^\w/, (c) => c.toUpperCase());
		}
	}
}
</script>

<style>
	#titlebar {
		display: flex;
		position: sticky;
		top: 0;
		height: 4rem;
		width: 100%;
		background-color: var(--white);
		opacity: 0.85;
	}
	#titlebar h4 {
		margin: 1rem;
		font-family: "Chirp Bold";
	}
	#titlebar #homeLink {
		font-size: 1.9rem;
		margin: 0.9rem 0 0 1rem;
		color: var(--black);
		border-radius: 50%;
	}
	#titlebar #homeLink:hover {
		background-color: var(--extraExtraLightGrey)
	}
</style>