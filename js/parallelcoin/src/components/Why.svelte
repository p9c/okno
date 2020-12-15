<script lang="js">
  import Shadow from '../../../components/txt/Shadow.svelte'
  import ParallelCoinLogo from '../../../components/svg/ParallelCoinLogo.svelte'
</script>


<div id="why" class="step slide" data-x="0" data-y="-3" data-scale="1" data-autoplay="0">
  <div class="w-slide h-slide flex flex-col justify-center items-center">
    <div class="w-full flex justify-end items-end">
      <div class="toplayer-lite bottomlayer-yellow">
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
        <img class="w-32 h-24" alt="" src="/square.png" />
      </div>
      <div class="toplayer-lite bottomlayer-orange">
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
        <img class="w-32 h-16" alt="" src="/square.png" />
      </div>
      <div class="toplayer-lite bottomlayer-red">
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
        <img class="w-32 h-8" alt="" src="/square.png" />
      </div>
      <div class="toplayer-lite bottomlayer-purple">
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
      </div>
      <div>
        <img class="w-32 h-8" alt="" src="/square.png" />
      </div>
    </div>
    <div class="p-10 font-brand uppercase text-5xl font-bold text-center">
      <Shadow>Why is ParallelCoin different than others?</Shadow>
    </div>
    <div class="w-full flex justify-start items-start">
      <div>
        <img class="w-32 h-8" alt="" src="/square.png" />
      </div>
      <div class="toplayer-lite bottomlayer-blue">
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
      </div>
      <div class="toplayer-lite bottomlayer-green">
        <img class="w-32 h-8" alt="" src="/square.png" />
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
      </div>
      <div class="toplayer-lite bottomlayer-gray">
        <img class="w-32 h-16" alt="" src="/square.png" />
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
      </div>
      <div class="toplayer-lite bottomlayer-dark">
        <img class="w-32 h-24" alt="" src="/square.png" />
        <ParallelCoinLogo styleClass={"w-32 h-32 p-3 svgFull"} />
      </div>
    </div>
  </div>
</div>
<style global>
.svgFull{
  max-width: 100%;
  height: auto;
  opacity: 0.6;
  transition: background 300ms ease-in-out;
}
.svgFull:hover {
  opacity: 0.6;
}

</style>
