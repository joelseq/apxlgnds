<script lang="ts">
  import type { Event } from '$lib/types'
  import { formatDuration, timeUntilEvent } from '$lib/utils'
  import { AccordionItem } from '@skeletonlabs/skeleton'

  const { event, withSeparator }: { event: Event; withSeparator: boolean } = $props()

  const timeClass =
    new Date(event.endDate).getTime() > new Date().getTime() ? 'text-green-600' : 'text-red-600'
</script>

<AccordionItem class="bg-gray-700 mx-auto p-4">
  <svelte:fragment slot="summary">
    <div class="flex justify-between">
      <p>{event.title}</p>
      <p class={timeClass}>{timeUntilEvent(event)}</p>
    </div>
  </svelte:fragment>
  <svelte:fragment slot="content">
    <div class="text-wrap break-words text-left space-y-2">
      <div>
        <strong>{formatDuration(event.startDate, event.endDate)}</strong>
      </div>
      {#if event.metadata?.battlefy_url}
        <div>
          <strong>Battlefy: </strong>
          <a
            class="underline mr-4"
            href={event.metadata.battlefy_url}
            target="_blank"
            title="View in battlefy"
          >
            {event.metadata.battlefy_url}
          </a>
        </div>
      {/if}
      {#if event.metadata?.reddit}
        <div>
          <strong>Reddit: </strong>
          <a class="underline mr-4" href={event.metadata.reddit.url}>
            {event.metadata.reddit.title}
          </a>
        </div>
      {/if}
    </div>
  </svelte:fragment>
</AccordionItem>
{#if withSeparator}
  <hr class="w-[400px] bg-gray-800 mx-auto" />
{/if}

<style>
</style>
