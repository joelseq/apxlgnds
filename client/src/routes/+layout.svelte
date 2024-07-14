<script>
  import '../app.css'
  import { dev } from '$app/environment'
  import { inject } from '@vercel/analytics'
  import { PlausibleAnalytics } from '@accuser/svelte-plausible-analytics'

  inject({ mode: dev ? 'development' : 'production' })
</script>

<PlausibleAnalytics domain="apxlgnds.com" apiHost="https://plausible.joelseq.app" />

<slot></slot>
