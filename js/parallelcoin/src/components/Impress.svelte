<script lang="js">
	import { onMount } from 'svelte';
	onMount(async () => {
		await import('impress.js');
		await import('impress.js/src/plugins/media/media');
		impress().init()
	});
</script>

<style global>
  @import '../assets/impress.css';
  @import '../../../components/css/svg.css';
</style>
