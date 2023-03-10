import { Ticker } from "../components/Ticker"


export const HighlightedTickers = () => {
  return (
    <>
    <div>Header</div>
    <div class="row"><Ticker symbol="CHR" /></div>
    <div class="row"><Ticker symbol="COG" /></div>
    </>
  )
}
