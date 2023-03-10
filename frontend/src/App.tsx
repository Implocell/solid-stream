import { Component } from 'solid-js';

import styles from './App.module.css';
import { Sliders } from './containers/Sliders';
import { HighlightedTickers } from './containers/HighlightedTickers';
import { Stocks } from './containers/Stocks';

const App: Component = () => {


  return (
    <div class={styles.App}>
     <HighlightedTickers />
     <Sliders />
     <Stocks />
    </div>
  );
};

export default App;
