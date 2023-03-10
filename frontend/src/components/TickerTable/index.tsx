import {
    Component,
    createEffect,
    createMemo,
    createSignal,
    For,
    onCleanup,
} from "solid-js";
import {
    subscribeToAllStocks,
    stocks,
    Ticker,
} from "../../data/stocks-with-nested-signals";
import styles from "./TickerTable.module.scss";

interface Sort {
    key: keyof Ticker;
    order: "ASC" | "DESC";
}

function getValueJSX(value: number): [digits: string, decimals: string] {
    return value.toFixed(2).split(".") as [string, string];
}

interface Props {
    initialSort?: Sort;
}

export const TickerTable: Component<Props> = ({ initialSort = { key: "value", order: "ASC" } }) => {
    // This logic could be separated, even if we need unique sort per instance
    const [sort, setSort] = createSignal<Sort>(initialSort);
    const [stocksCopy, setStocksCopy] = createSignal(stocks());

    // If we found a case where this was slow, we could look into Store or createMutable perhaps
    createEffect(() => {
        const { key, order } = sort();

        setStocksCopy(() => {
            return [
                ...stocks().sort((a, b) => {
                    if (order === "ASC") {
                        return (
                            (a.ticker()[key] as number) -
                            (b.ticker()[key] as number)
                        );
                    }

                    return (
                        (b.ticker()[key] as number) -
                        (a.ticker()[key] as number)
                    );
                }),
            ];
        });
    });

    const onHeaderClick = (key: string) => {
        setSort((prev) => {
            const isSame = prev.key === key;

            if (isSame) {
                return {
                    key,
                    order: prev.order === "ASC" ? "DESC" : "ASC",
                } as Sort;
            }

            return { key, order: "ASC" } as Sort;
        });
    };

    const clearSub = subscribeToAllStocks();

    onCleanup(() => clearSub());

    return (
        <div class={styles.table} role="table">
            <div class={styles.row} role="rowheader">
                <div
                    class={`${styles.cell} ${styles.name}`}
                    role="columnheader"
                    onClick={() => onHeaderClick("name")}
                >
                    Name
                </div>
                <div
                    class={`${styles.cell} ${styles.value}`}
                    role="columnheader"
                    onClick={() => onHeaderClick("value")}
                >
                    Last
                </div>
                <div
                    class={`${styles.cell} ${styles.updated}`}
                    role="columnheader"
                    onClick={() => onHeaderClick("updated")}
                >
                    Updated
                </div>
            </div>
            <For each={stocksCopy()}>
                {({ ticker }) => {
                    const valueAccessor = createMemo(() =>
                        getValueJSX(ticker().value)
                    );

                    return (
                        <div class={styles.row} role="row">
                            <div
                                class={`${styles.cell} ${styles.name}`}
                                role="cell"
                            >
                                [{ticker().symbol}] {ticker().name}
                            </div>
                            <div
                                class={`${styles.cell} ${styles.value}`}
                                role="cell"
                            >
                                <span class={styles.digit}>
                                    {valueAccessor()[0]}.
                                </span>
                                <span class={styles.decimals}>
                                    {valueAccessor()[1]}
                                </span>
                            </div>
                            <div
                                class={`${styles.cell} ${styles.updated}`}
                                role="cell"
                            >
                                {new Date(
                                    ticker().updated
                                ).toLocaleTimeString()}
                            </div>
                        </div>
                    );
                }}
            </For>
        </div>
    );
};
