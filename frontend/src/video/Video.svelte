<script>
	import VideoPlayer from './VideoPlayer.svelte'
	import VideoSeekbar from './VideoSeekbar.svelte'
	import VideoControls from './VideoControls.svelte'
</script>

<div class="video">
	<VideoPlayer />
	<VideoSeekbar />
	<VideoControls />
</div>

<style>
	.video {
		flex: 1 0 100px;
		height: 100%;

		display: flex;
		flex-direction: column;

		& > :global(:first-child) {
			flex: 1 0 400px;
		}

		& > :global(:not(:first-child)) {
			border-top: 1px solid black;
		}
	}
</style>
