import { TickerTable } from "../../components/TickerTable";
import styles from './Stocks.module.scss';

export const Stocks = () => {
    return (
        <>
            <div>Header</div>
            <div class={styles.stocksContainer}>
                <TickerTable />
                <TickerTable />
            </div>
        </>
    );
};
