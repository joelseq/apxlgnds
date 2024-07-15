<script lang="ts">
  import type { Event } from '$lib/types'
  import type { PageData } from './$types'
  import EventCard from '$lib/components/EventCard.svelte'
  import { Accordion } from '@skeletonlabs/skeleton'

  export let data: PageData

  const { algs, other } = data.events
</script>

<div class="flex items-center justify-center">
  <div class="container p-4 text-center mb-24">
    <h1 class="text-6xl logo-text my-12">APXLGNDS</h1>
    <!-- <h2 class="text-2xl my-4">Links</h2> -->
    <div class="flex gap-12 justify-center items-center">
      <a href="https://battlefy.com/apex-legends-global-series-year-4" target="_blank">
        <img src="/logo-battlefy-white.svg" alt="Battlefy Logo" width={125} />
      </a>
      <a href="https://liquipedia.net/apexlegends/Main_Page" target="_blank">
        <img src="/liquipedia_icon_menu.png" alt="Liquipedia Logo" height={50} />
      </a>
      <a href="https://reddit.com/r/CompetitiveApex" target="_blank">
        <img src="/reddit-logo.svg" alt="Reddit Logo" width={50} />
      </a>
    </div>
    <div class="flex flex-wrap justify-center gap-12">
      <div class="section">
        <h2 class="text-2xl my-4">Upcoming Events</h2>
        <div class="mx-auto">
          {@render eventContainer(algs.upcoming, 'ALGS events')}
          {@render eventContainer(other.upcoming, 'Other events')}
        </div>
      </div>
      <div class="section">
        <h2 class="text-2xl my-4">Recent Events</h2>
        <div class="mx-auto">
          {@render eventContainer(algs.recent, 'ALGS events')}
          {@render eventContainer(other.recent, 'Other events')}
        </div>
      </div>
    </div>
  </div>
</div>

{#snippet eventContainer(events: Event[], title)}
  <div class="mb-4">
    <h3 class="text-xl bg-gray-700 pt-4">
      {title}
    </h3>
    {#if events == null || events.length < 1}
      <div class="bg-gray-700 p-4">
        <p>No events to show</p>
      </div>
    {:else}
      <Accordion hover="hover:bg-none" spacing="space-y-0">
        {#each events as event, i}
          <EventCard {event} withSeparator={i !== events.length - 1} />
        {/each}
      </Accordion>
    {/if}
  </div>
{/snippet}

<style lang="postcss">
  .logo-text {
    color: #68160f;
    font-family: 'Anton SC', sans-serif;
    font-style: normal;
  }

  .section {
    @apply text-center sm:w-3/4 lg:w-2/5;
  }
</style>
