<script>
  import LogIn from './LogIn.svelte';
  import Register from './Register.svelte';
</script>

<style>
  .main {
    align-items: center;
    background: linear-gradient(90deg, #515093 0%, rgb(90, 111, 180) 100%);
    display: flex;
    height: 100vh;
    justify-content: center;
  }

  .container {
    background: white;
    width: 480px;
  }

  .title {
    font-family: Quicksand, sans-serif;
    font-size: 3rem;
    font-weight: normal;
    margin: 24px 0;
    text-align: center;
  }
</style>

<main class="main">
  <div class="container">
    <h1 class="title">OKNO</h1>
    <div><LogIn></LogIn></div>
    <hr>
    <div><Register></Register></div>
  </div>
</main>
