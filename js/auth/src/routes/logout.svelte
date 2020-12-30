<script context="module">
  export async function preload (page, session) {
    await this.fetch('http://a.okno.rs:4433/authenticate', {
      method: 'DELETE',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
  }
</script>

<script>
  import { goto, stores } from '@sapper/app'

  const { session } = stores()

  session.set({ authenticated: false, profile: null })

  goto('/')
</script>
