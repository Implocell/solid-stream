import { createEffect, createSignal, onCleanup } from "solid-js"


export const Ticker = ({ticker}:{ticker:string}) => {
    const [value, setValue] = createSignal("")
    const ws = new WebSocket(`ws://localhost:4503/updates/${ticker}`)

    ws.onmessage = (e) => {
      setValue(`${e.data}`)
    }
    onCleanup(()=> ws.close())
  return (
    <div>
        <span>Current Time:</span>{value()}</div>
  )
}
