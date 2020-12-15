<script>
  import Tailwind from '../components/Tailwind.svelte';
  import Impress from '../components/Impress.svelte';
  import Intro from '../components/Intro.svelte';
  import What from '../components/What.svelte';
  import Why from '../components/Why.svelte';
  import Fun from '../components/Fun.svelte';
  import Step from '../components/Step.svelte';
  import Download from '../components/Download.svelte';
  import Doc from '../components/Doc.svelte';
  import Story from '../components/Story.svelte';
  import Community from '../components/Community.svelte';
  import Status from '../components/Status.svelte';
  import Explorer from '../components/Explorer.svelte';
  import Borg from '../components/Borg.svelte';
  import Where from '../components/Where.svelte';
  import Video from '../components/Video.svelte';

</script>

<svelte:head>
	<title>Let's go Parallel</title>
</svelte:head>

<Video />
<Tailwind />
<Impress />




<div id="impress"
    data-transition-duration="1000"

    data-width="1200"
    data-height="900"
    data-max-scale="3"
    data-min-scale="0"
    data-perspective="0"

    data-media-autoplay="true"
    data-media-autostop="true"
>

  <Intro />
  <What />
  <Why />
  <Fun />
  <Step />
  <Download />
  <Story />
  <Doc />
  <Community />
  <Status />
  <Explorer />
  <Borg />
  <Where />

</div>

<style global>
.transparentText{
  opacity: 0.62;
}
</style>
