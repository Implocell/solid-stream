import { Accessor, createSignal, Setter } from "solid-js";

export interface Ticker {
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

let messagesSincePrint = 0;
let intervalId: NodeJS.Timer;

const close = () => {
    clearInterval(intervalId);
    if (--subs === 0 && ws) {
        ws.close();
    }
};

export function subscribeToAllStocks() {
    if (ws) {
        return close;
    }

    intervalId = setInterval(() => {
        console.log(
            `subscribeToAllStocks: ${messagesSincePrint} messages last 10 seconds (${
                messagesSincePrint / 10
            } msg/s)`
        );
        messagesSincePrint = 0;
    }, 10 * 1000);

    ws = new WebSocket(`ws://localhost:4503/stocks`);

    ws.onmessage = (e) => {
        const next = JSON.parse(e.data) as Ticker;

        setTickers((tickers) => {
            const foundTicker = tickers.find(
                (current) => current.ticker().symbol === next.symbol
            );

            messagesSincePrint++;

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
