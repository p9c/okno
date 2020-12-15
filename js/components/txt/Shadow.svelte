<script>
  export let styleClass;
</script>

  <h2 class='artificial'><slot/></h2>
  <!--  <h2 class='ambient'><slot/></h2>-->


<style>
  h2 {
    line-height: 100%;
    font-family: plan9regular, "Avant Garde", Avantgarde, "Century Gothic", CenturyGothic, "AppleGothic", sans-serif;
    font-size: 100pt;
    text-align: center;
    text-transform: uppercase;
    text-rendering: optimizeLegibility;
    display: block;
    float: left;
    padding-bottom: 100pt;
  }
  h2.artificial {
    color: var(--lite-color);
    letter-spacing: .15em;
    text-shadow: 3px -6px 0px rgba(0,0,0,0.1), 2px -4px 0px rgba(0,0,0,0.1), 1px -2px 0px rgba(0,0,0,0.1), 0px 0px 0px rgba(0,0,0,0.1), -1px 2px rgba(0,0,0,0.2), -2px 4px rgba(0,0,0,0.2), -3px 6px rgba(0,0,0,0.2), -4px 8px rgba(0,0,0,0.2), -5px 10px rgba(0,0,0,0.2), -6px 12px 1px rgba(0,0,0,0.2), -7px 14px 1px rgba(0,0,0,0.2), -8px 16px 1px rgba(0,0,0,0.2), -9px 18px 1px rgba(0,0,0,0.2), -10px 20px 1px rgba(0,0,0,0.2), -11px 22px 1px rgba(0,0,0,0.2), -12px 24px 1px rgba(0,0,0,0.2), -13px 26px 1px rgba(0,0,0,0.2), -14px 28px 1px rgba(0,0,0,0.2), -15px 30px 1px rgba(0,0,0,0.2), -16px 32px 1px rgba(0,0,0,0.2), -17px 34px 1px rgba(0,0,0,0.2), -18px 36px 1px rgba(0,0,0,0.2), -19px 38px 1px rgba(0,0,0,0.2), -20px 40px 1px rgba(0,0,0,0.2);
  }
  h2.ambient{
    color: #131313;
    letter-spacing: .15em;
    text-shadow: 3px -6px 0px rgba(0,0,0,0.1), 2px -4px 0px rgba(0,0,0,0.1), 1px -2px 0px rgba(0,0,0,0.1), 0px 0px 0px rgba(0,0,0,0.1), -1px 2px 1px rgba(0,0,0,0.3), -2px 4px 1px rgba(0,0,0,0.3), -3px 6px 1px rgba(0,0,0,0.3), -4px 8px 1px rgba(0,0,0,0.3), -5px 10px 1px rgba(0,0,0,0.3), -6px 12px 1px rgba(0,0,0,0.3), -7px 14px 1px rgba(0,0,0,0.3), -8px 16px 1px rgba(0,0,0,0.3), -9px 18px 1px rgba(0,0,0,0.3), -10px 20px 1px rgba(0,0,0,0.3), -11px 22px 1px rgba(0,0,0,0.3), -12px 24px 1px rgba(0,0,0,0.3), -13px 26px 1px rgba(0,0,0,0.3), -14px 28px 1px rgba(0,0,0,0.3), -15px 30px 1px rgba(0,0,0,0.3), -16px 32px 1px rgba(0,0,0,0.3), -17px 34px 1px rgba(0,0,0,0.3), -18px 36px 1px rgba(0,0,0,0.3), -19px 38px 1px rgba(0,0,0,0.3), -20px 40px 1px rgba(0,0,0,0.3);
  }
</style>
