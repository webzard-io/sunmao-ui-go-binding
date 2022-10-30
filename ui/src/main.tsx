import { implementUtilMethod } from "@sunmao-ui/runtime";
import React from "react";
import ReactDOM from "react-dom";
import App from "./App";

type Options = {
  application: unknown;
  modules: any[];
  wsUrl: string;
  reloadWhenWsDisconnected: boolean;
  handlers: string[];
  utilMethods?: { options: any; impl: any }[];
};

export function renderApp(options: Options) {
  const {
    wsUrl,
    application,
    modules,
    reloadWhenWsDisconnected,
    handlers,
    utilMethods,
  } = options;
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
        modules={modules}
        ws={ws}
        handlers={handlers}
        utilMethods={utilMethods?.map(
          (u) => () => implementUtilMethod(u.options)(u.impl)
        )}
      />
    </React.StrictMode>,
    document.getElementById("root")!
  );
}
