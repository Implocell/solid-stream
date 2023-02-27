import { Ticker } from "../components/Ticker"


export const Tickers = () => {
  return (
    <>
    <div>Header</div>
    <div class="row"><Ticker symbol="CHR" /></div>
    <div class="row"><Ticker symbol="COG" /></div>
    </>
  )
}
