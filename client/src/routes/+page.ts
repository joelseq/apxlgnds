import type { EventsResponse } from '$lib/types'
import axios from 'axios'

export async function load() {
  const response = await axios.get('http://localhost:8080/events')

  const events: EventsResponse['events'] = response.data.events

  return {
    events,
  }
}
