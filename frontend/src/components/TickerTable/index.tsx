import { For, onCleanup } from "solid-js";
import { subscribeToAllStocks, stocks } from "../../data/stocks-with-nested-signals";
import styles from './TickerTable.module.scss';

export const TickerTable = () => {
    const clearSub = subscribeToAllStocks();

    onCleanup(() => clearSub());

    return (
        <div class={styles.table} role="table">
            <div class={styles.row} role="rowheader">
                <div class={`${styles.cell} ${styles.name}`} role="columnheader">
                    Name
                </div>
                <div class={`${styles.cell} ${styles.value}`} role="columnheader">
                    Last
                </div>
                <div class={`${styles.cell} ${styles.updated}`} role="columnheader">
                    Updated
                </div>
            </div>
            <For each={stocks()}>
                {({ ticker }) => (
                    <div class={styles.row} role="row">
                        <div class={`${styles.cell} ${styles.name}`} role="cell">
                            [{ticker().symbol}] {ticker().name}
                        </div>
                        <div class={`${styles.cell} ${styles.value}`} role="cell">
                            {ticker().value?.toFixed(2)}
                        </div>
                        <div class={`${styles.cell} ${styles.updated}`} role="cell">
                            {new Date(ticker().updated).toLocaleTimeString()}
                        </div>
                    </div>
                )}
            </For>
        </div>
    );
};
