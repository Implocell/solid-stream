import { Accessor, createSignal, Setter } from "solid-js";

interface Ticker {
    symbol: string;
    name: string;
    value: number;
    updated: number;
}

type TickerListEntry = {
    ticker: Accessor<Ticker>;
    setTicker: Setter<Ticker>;
};

const [tickersAccessor, setTickers] = createSignal<TickerListEntry[]>([]);

export const stocks = tickersAccessor;

let ws: WebSocket;
let subs = 0;

const close = () => {
    if (--subs === 0 && ws) {
        ws.close();
    }
};

export function subscribeToAllStocks() {
    if (ws) {
        return close;
    }

    ws = new WebSocket(`ws://localhost:4503/stocks`)

    ws.onmessage = (e) => {
        const next = JSON.parse(e.data) as Ticker;

        setTickers((tickers) => {
            const foundTicker = tickers.find(
                (current) => current.ticker().symbol === next.symbol
            );

            if (foundTicker) {
                foundTicker.setTicker(next);

                return tickers;
            }

            const [ticker, setTicker] = createSignal(next);

            return [...tickers, { ticker, setTicker }];
        });
    };

    return close;
}