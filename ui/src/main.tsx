import React from "react";
import ReactDOM from "react-dom";
import App from "./App";

type Options = {
  application: unknown;
  wsUrl: string;
  reloadWhenWsDisconnected: boolean;
  handlers: string[];
};

export function renderApp(options: Options) {
  const { wsUrl, application, reloadWhenWsDisconnected, handlers } = options;
  const ws = new WebSocket(wsUrl);
  ws.onopen = () => {
    console.log("ws connected");
  };
  ws.onclose = () => {
    if (reloadWhenWsDisconnected) {
      setTimeout(() => {
        window.location.reload();
      }, 1500);
    }
  };

  ReactDOM.render(
    <React.StrictMode>
      <App
        application={application}
        onStoreChange={() => {}}
        ws={ws}
        handlers={handlers}
      />
    </React.StrictMode>,
    document.getElementById("root")!
  );
}
