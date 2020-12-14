<script>
  import notify from '../helpers/notify';

  import { jwt } from '../stores';

  let user = '';
  let password = '';

  async function logIn(e) {
    e.preventDefault();
    const req = { user, password };
    const res = await fetch('http://api.okno.rs:4433/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    });
    if (res.ok) {
      notify('Logged in', 'green');
      const json = await res.json();
      localStorage.setItem('jwt', json.token);
      jwt.set(json.token);
    } else {
      notify('Could not log in', 'red');
    }
  }
</script>

<style>
  .title {
    font-family: Quicksand, sans-serif;
    font-size: 2rem;
    font-weight: normal;
    margin: 16px 0;
    text-align: center;
  }
</style>


    <h3 class="title">Log In</h3>
    <form class="form" on:submit="{logIn}">
      <label class="field">
        <span class="label">User</span>
        <input class="input" bind:value="{user}" />
      </label>
      <label class="field">
        <span class="label">Password</span>
        <input type="password" class="input" bind:value="{password}" />
      </label>
      <button type="submit" class="button button-primary">Log in</button>
    </form>
