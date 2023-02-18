import { Ticker } from "../components/Ticker"


export const Tickers = () => {
  return (
    <>
    <div>Header</div>
    <div class="row"><Ticker ticker="first" /></div>
    <div class="row"><Ticker ticker="second" /></div>
    </>
  )
}
