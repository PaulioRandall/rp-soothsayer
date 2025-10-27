<script context="module">
	import Two from 'two.js'

	function createTwoInstance(container) {
		if (!container) {
			return null
		}

		return new Two({
			fullscreen: true,
			autostart: true,
		}).appendTo(container)
	}

	function addContentToTwoInstance(two, dirFiles) {
		if (!two || !dirFiles) {
			return
		}

		const size = 50
		const radius = size / 2
		const gap = 50
		const childY = size + gap + radius

		two.makeRectangle(two.width / 2, radius, size, size)
		two.makeText('Parent', two.width / 2, radius)

		for (let i = 0; i < dirFiles.length; i++) {
			const dir = dirFiles[i]
			const x = i * (gap + size) + radius
			two.makeRectangle(x, childY, size, size)
			two.makeText(dir.DirName, x, childY)
		}
	}
</script>

<script>
	//import { ReadDir } from '../../wailsjs/go/api/App'
	import MockDirInfoData from './MockDirInfoData'

	let dirFiles = $state(MockDirInfoData)
	let dirFileCount = $derived(dirFiles.length)

	let vizContainer = $state()
	let twoInstance = null

	$effect(() => {
		twoInstance = createTwoInstance(vizContainer)
		addContentToTwoInstance(twoInstance, dirFiles)
	})

	//dirFiles = ReadDir(".")
	//console.log(dirFiles)
</script>

<main>
	<div bind:this={vizContainer} class="file-browser-viz-container">
		<!-- Inner HTML managed by Two.js -->
	</div>
</main>

<style>
	main {
		width: 100%;
		height: 100%;
	}

	.file-browser-viz-container {
		width: 100%;
		height: 100%;
		max-width: 100%;
		max-height: 100%;
	}
</style>
