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

<Video/>
<Tailwind/>
<Impress/>




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

  <Intro></Intro>
  <What></What>
  <Why></Why>
  <Fun></Fun>
  <Step></Step>
  <Download></Download>



  <Story></Story>
  <Doc></Doc>
  <Community></Community>

  <Status></Status>
  <Explorer></Explorer>


  <Borg></Borg>
  <Where></Where>

</div>








<!--


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
  <div id="intro" class="step slide" data-x="-140" data-y="0" data-scale="10" data-rotate-x="90" _data-autoplay="10">
    <div class="w-slide h-slide flex flex-col justify-center items-center">
      <div class="w-full flex justify-center items-center">
        <div>
          <img class="w-72 h-72" alt="" src="/sailsonic_logo_contour.svg" />
        </div>
      </div>
      <div class="p-10 font-brand uppercase text-5xl font-bold text-center">
        Let's go windsurfing!
      </div>
    </div>
  </div>

  <div id="what" class="step slide" data-x="0" data-y="0" data-scale="1" data-rotate-x="90" data-rotate-z="90" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        What is windsurfing?
      </div>
      <div>
      </div>
    </div>
  </div>

  <div id="why" class="step slide" data-x="0" data-y="-3" data-scale="1" data-autoplay="2">
    <div class="w-slide h-slide flex flex-col justify-center items-center">
      <div class="w-full flex justify-end items-end">
        <div>
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
          <img class="w-32 h-24" alt="" src="/black-square.jpg" />
        </div>
        <div>
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
          <img class="w-32 h-16" alt="" src="/black-square.jpg" />
        </div>
        <div>
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
          <img class="w-32 h-8" alt="" src="/black-square.jpg" />
        </div>
        <div>
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
        </div>
        <div>
          <img class="w-32 h-8" alt="" src="/black-square.jpg" />
        </div>
      </div>
      <div class="p-10 font-brand uppercase text-5xl font-bold text-center">
        Why is Windsurfing good for you?
      </div>
      <div class="w-full flex justify-start items-start">
        <div>
          <img class="w-32 h-8" alt="" src="/black-square.jpg" />
        </div>
        <div>
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
        </div>
        <div>
          <img class="w-32 h-8" alt="" src="/black-square.jpg" />
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
        </div>
        <div>
          <img class="w-32 h-16" alt="" src="/black-square.jpg" />
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
        </div>
        <div>
          <img class="w-32 h-24" alt="" src="/black-square.jpg" />
          <img class="w-32 h-32" alt="" src="/sailsonic_logo_contour.svg" />
        </div>
      </div>
    </div>
  </div>

  <div id="it-is-fun" class="step slide" data-x="36" data-y="-196" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        Absolute fun. Amazing sensations. Healthy addiction.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-challenge" class="step slide" data-x="-12" data-y="256" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        Easy to start. Hard to advance.<br/>Never gets boring.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-safe" class="step slide" data-x="164" data-y="-160" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        It is very safe.<br/>Need for an adrenaline? Go extreme.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-healthy" class="step slide" data-x="-144" data-y="224" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        Improves physical and mental health.<br/>Full body exercise. Intense energy release.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-environmentally-friendly" class="step slide" data-x="292" data-y="-128" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-center">
        It is environmentally friendly.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-affordable" class="step slide" data-x="-272" data-y="188" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        It is relatively affordable sport.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-olympic" class="step slide" data-x="420" data-y="-94" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        Lots of options. And it is Olympic sport.
      </div>
      <div>
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="it-is-lifetime" class="step slide" data-x="-400" data-y="152" data-scale="0.05" _data-autoplay="10">
    <div class="w-slide h-slide  flex flex-col justify-center items-center">
      <div class="p-10 w-full font-brand uppercase text-4xl font-bold text-left">
        It is Lifetime sport.
      </div>
    </div>
  </div>
  <div class="step slide" data-x="0" data-y="0" data-scale="0.8" data-autoplay="0.001" />

  <div id="welcome" class="step slide" data-x="-140" data-z="-5000" data-scale="6" data-rotate-x="90" data-autoplay="2">
    <div class="w-slide h-slide flex flex-col justify-center items-center">
      <div class="p-10 font-brand uppercase text-5xl font-bold text-center">
        And you can windsurf too!
      </div>
    </div>
  </div>

  <div id="where" class="step slide" data-x="-140" data-z="10000" data-scale="6" data-rotate-x="90" _data-autoplay="10">
    <div class="w-slide h-slide flex flex-col justify-start items-center">
      <div class="p-10 font-brand uppercase text-5xl font-bold text-center">
        <span class="">Naval Point Sailing Club</span>
      </div>
      <div class="pt-4 prose-2xl text-white-default">
      </div>
    </div>
  </div>

</div>
-->


<style global>
.transparentText{
  opacity: 0.62;
}
</style>
