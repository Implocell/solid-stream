import { createEffect, createSignal, onCleanup } from "solid-js"

interface Props {
  symbol: string
}

interface Ticker {
  symbol: string;
  name: string;
  value: number;
  updated: number;
}

export const Ticker = ({ symbol }: Props) => {
    const [ticker, setTicker] = createSignal({} as Ticker)

    const ws = new WebSocket(`ws://localhost:4503/stock/${symbol}`)

    ws.onmessage = (e) => {
      setTicker(JSON.parse(e.data));
    }

    onCleanup(()=> ws.close())
  return (
    <div>
      <h3>[{ticker().symbol}] {ticker().name}</h3>
      <span>{ticker().value?.toFixed(2)} @ {(new Date(ticker().updated)).toLocaleTimeString()}</span>
    </div>
  )
}
