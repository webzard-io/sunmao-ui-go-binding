import { implementUtilMethod } from "@sunmao-ui/runtime";
import React from "react";
import ReactDOM from "react-dom";
import Editor from "./Editor";
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
      <Editor
        application={application}
        modules={modules}
        ws={ws}
        handlers={handlers}
        utilMethods={utilMethods?.map(
          (u) => () => implementUtilMethod(u.options)(u.impl)
        )}
        applicationPatch={applicationPatch}
        modulesPatch={modulesPatch}
      />
    </React.StrictMode>,
    document.getElementById("root")!
  );
}
