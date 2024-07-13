import { PUBLIC_BASE_API_URL } from '$env/static/public'
import type { EventsResponse } from '$lib/types'
import axios from 'axios'

export async function load() {
  const response = await axios.get(`${PUBLIC_BASE_API_URL}/events`)

  const events: EventsResponse = response.data

  return {
    events,
  }
}
