<template>
	<div class="notification">
		<div class="header">
			<h6>{{ sender.firstName }} {{ sender.lastName }}</h6>
		</div>
		
		<div class="body">
			<p>{{ message.content }}</p>
		</div>
	</div>
</template>

<script>
export default {
	name: 'Template_Notification',
	props: {
		message: Object,
		sender: Object,
	},
}
</script>
