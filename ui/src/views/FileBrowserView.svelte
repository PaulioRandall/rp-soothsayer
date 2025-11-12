<script>
	import {
		ParentPath,
		AbsPath,
		ReadDir,
		CreateStudy,
	} from '../../wailsjs/go/api/App'
	import { Button, TextInput } from '../lib'

	let dirFiles = $state([])
	let absDirPath = $state('')
	let backPath = $state('')
	let dirPath = $state('.')

	let showProjectNameInput = $state(false)
	let newProjectName = $state('')

	$effect(async () => {
		absDirPath = await AbsPath(dirPath)
		backPath = await ParentPath(absDirPath)
		dirFiles = await ReadDir(absDirPath)
	})

	async function createNewProject() {
		const name = newProjectName.trim()
		const path = absDirPath

		if (!name) {
			return
		}

		showProjectNameInput = false
		await CreateStudy(name, path)
	}

	const showProjectNameBox = () => (showProjectNameInput = true)
	const hideProjectNameBox = () => (showProjectNameInput = false)
	const gotoParentDir = () => (dirPath = backPath)
</script>

<main>
	<div class="file-browser-view">
		<div class="dir-info">
			<span>
				<Button onclick={gotoParentDir}>Back</Button>
			</span>
			<span class="dir-path">
				{absDirPath}
			</span>
			<span>
				<Button onclick={showProjectNameBox}>Create Project Here</Button>
			</span>
		</div>
		<div class="dir-files">
			{#each dirFiles as { DirPath, DirName, IsProjectDir }}
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
							<Button onclick={() => (dirPath = DirPath)}>➜</Button>
						{/if}
					</span>
				</dir>
			{/each}
		</div>
	</div>

	{#if showProjectNameInput}
		<div class="project-name-modal-bg" onclick={hideProjectNameBox}>
			<div class="project-name-modal" onclick={(e) => e.stopPropagation()}>
				<label class="project-name-input-label" for="project-name-input"
					>Project Name</label>
				<TextInput
					autofocus
					id="project-name-input"
					bind:value={newProjectName} />
				<Button class="submit-project-name-button" onclick={createNewProject}
					>Create</Button>
			</div>
		</div>
	{/if}
</main>

<style>
	main {
		position: relative;

		width: 100%;
		height: 100%;
	}

	.file-browser-view {
		width: 100%;
		max-width: 100%;
		height: 100%;
		max-height: 100%;

		display: flex;
		flex-direction: column;
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
		flex: 1 1 auto;
		overflow: auto;
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

	.project-name-modal-bg {
		position: absolute;
		display: flex;
		justify-content: center;
		align-items: center;

		top: 0;
		left: 0;

		width: 100%;
		height: 100%;

		background: #88888888;
	}

	.project-name-input-label {
		font-weight: bold;
		width: 100%;
	}

	.project-name-modal {
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		gap: 0.5rem;

		max-width: 600px;
		max-height: 400px;

		padding: 1rem;

		border: 2px solid black;
		border-radius: 6px;
		background: white;
	}

	.submit-project-name-button {
		align-self: flex-end;
	}
</style>
