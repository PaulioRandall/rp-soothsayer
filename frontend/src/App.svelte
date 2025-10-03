<script>
	import { Video } from './video'
	import { Notes } from './notes'
	import AdjustableLayout from './AdjustableLayout.svelte'
</script>

<main>
	<div class="main-layout">
		<AdjustableLayout left={Video} right={Notes} />
	</div>
</main>

<style>
	main {
		margin: 0;
		padding: 0;

		width: 100vw;
		height: 100vh;
	}

	.main-layout {
		height: 100%;
		width: 100%;

		max-height: 100%;
		max-width: 100%;

		overflow: hidden;
	}
</style>
