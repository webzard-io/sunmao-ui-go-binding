import { implementUtilMethod } from "@sunmao-ui/runtime";
import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { MainOptions } from "./shared";

export function renderApp(options: MainOptions) {
  const {
    wsUrl,
    application,
    modules,
    reloadWhenWsDisconnected,
    handlers,
    utilMethods,
    applicationPatch,
    modulesPatch,
    applicationBase
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
        applicationPatch={applicationPatch}
        modulesPatch={modulesPatch}
        applicationBase={applicationBase}
      />
    </React.StrictMode>,
    document.getElementById("root")!
  );
}
