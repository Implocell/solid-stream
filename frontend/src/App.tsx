import { Component } from 'solid-js';

import styles from './App.module.css';
import { Sliders } from './containers/Sliders';
import { Tickers } from './containers/Tickers';

const App: Component = () => {


  return (
    <div class={styles.App}>
     <Tickers />
     <Sliders />
    </div>
  );
};

export default App;
