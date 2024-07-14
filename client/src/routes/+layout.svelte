<script>
  import '../app.css'
  import { dev } from '$app/environment'
  import { inject } from '@vercel/analytics'
  import { PlausibleAnalytics } from '@accuser/svelte-plausible-analytics'

  inject({ mode: dev ? 'development' : 'production' })
</script>

<PlausibleAnalytics domain="apxlgnds.com" apiHost="https://plausible.joelseq.app" />

<slot></slot>

<footer class="w-full text-center py-5 bg-slate-800">
  <p class="text-gray-400">
    Created by <a class="text-blue-500" href="https://x.com/@_joelseq" target="_blank"
      >Joel Sequeira</a
    >
    | <a href="https://github.com/joelseq/apxlgnds">GitHub</a>
  </p>
</footer>

<style>
  footer a {
    text-decoration: none;
  }
  footer a:hover {
    text-decoration: underline;
  }
</style>
