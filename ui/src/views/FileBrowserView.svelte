<script>
	//import { ReadDir } from '../../wailsjs/go/api/App'
	import MockDirInfoData from './MockDirInfoData'

	let dirFiles = $state(MockDirInfoData) 

	//dirFiles = ReadDir(".")
	//console.log(dirFiles)
</script>

<main>
	{#each dirFiles as { ParentPath, DirName, IsProjectDir } }
		<div>
			{DirName}
		</div>
	{/each}
</main>
