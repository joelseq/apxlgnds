export type Event = {
  title: string
  description: string
  startDate: string
  endDate: string
}

export type EventsResponse = {
  events: Event[]
}
