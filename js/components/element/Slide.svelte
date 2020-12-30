<script lang="js">
  import Shadow from './../txt/Shadow.svelte'
  export let idTag;
  export let styleClass;
  export let dataX;
  export let dataY;
  export let dataScale;
  export let dataAuto;
  export let slideTitle;
</script>

<div id={idTag} class={styleClass} data-x={dataX} data-y={dataY} data-scale={dataScale} _data-autoplay={dataAuto}>
  <div class="w-slide h-slide  flex flex-col justify-center items-center">
    <div class="font-brand uppercase text-left">
      <Shadow styleClass={"artificial"}>{slideTitle}</Shadow>
    </div>
    <div class="p-10 w-full text-2xl text-left"><slot /></div>
  </div>
</div>
<style global>

</style>
