<script>
	import { AbsPath, ReadDir } from '../../wailsjs/go/api/App'
	import { Button } from '../lib'

	let dirFiles = $state([])
	let dirPath = $state('./')
	let absPath = $state('')

	$effect(async () => {
		absPath = await AbsPath(dirPath)
		dirFiles = await ReadDir(absPath)
	})
</script>

<main>
	<div class="file-browser-view">
		<div class="dir-info">
			<span>
				<Button>Back</Button>
			</span>
			<span class="dir-path">
				{absPath}
			</span>
			<span>
				<Button>Create Project Here</Button>
			</span>
		</div>
		<div class="dir-files">
			<!-- TODO render list of folders in the current dir path -->
			{#each dirFiles as { ParentPath, DirName, IsProjectDir }}
				<dir class="dir-file">
					<span class="dir-file-name">
						{#if IsProjectDir}
							*
						{/if}
						/{DirName}
					</span>
					<span>
						{#if IsProjectDir}
							<Button>Open Project</Button>
						{:else}
							<Button>➜</Button>
						{/if}
					</span>
				</dir>
			{/each}
		</div>
	</div>
</main>

<style>
	main {
		width: 100%;
		height: 100%;
	}

	.file-browser-view {
		width: 100%;
		max-width: 100%;
		min-height: 100%;
	}

	.dir-info {
		display: flex;

		padding: 1rem;

		border-bottom: 2px solid black;
	}

	.dir-path {
		flex: 1 1 auto;

		display: flex;
		align-items: center;
		padding: 0 1rem;

		font-size: 120%;
	}

	.dir-files {
		display: flex;
		flex-direction: column;

		overflow-y: auto;
	}

	.dir-file {
		display: flex;
		justify-content: space-between;

		padding: 0.25rem 1rem;

		border: 1px solid lightgrey;

		margin: 0;
	}

	.dir-file-name {
		display: flex;
		align-items: center;
	}
</style>
