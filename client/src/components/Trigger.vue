<template>
	<span ref="trigger"></span>
</template>

<script>
	export default {
		name: 'TriggerIntersect',
		props: {
			options: {
				type: Object,
				default() {
					return {
						root: null,
						threshold: "0",
					}
				}
			}
		},
		data() {
			return {
				observer: null,
			}
		},
		mounted() {
			this.observer = new IntersectionObserver( entries => {
				this.handleIntersect(entries[0])
			}, this.options)

			this.observer.observe(this.$refs.trigger)
		},
		unmounted() {
			this.observer.disconnect()
		},
		methods: {
			handleIntersect(entry) {
				if (entry.isIntersecting) this.$emit("triggerIntersected")
			}
		}
	}
</script>