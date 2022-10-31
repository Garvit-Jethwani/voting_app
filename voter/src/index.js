import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import { asyncWithLDProvider } from 'launchdarkly-react-client-sdk';

const renderApp = async () => {

  const LDProvider = await asyncWithLDProvider({ clientSideID: '635f5f4e009f201200f279b3' });
ReactDOM.render(
  <React.StrictMode>
    <LDProvider>
    <App />
    </LDProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
}
renderApp();
