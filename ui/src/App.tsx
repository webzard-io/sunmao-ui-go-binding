import { useEffect } from "react";
import { implementUtilMethod, initSunmaoUI, watch } from "@sunmao-ui/runtime";
import { sunmaoChakraUILib } from "@sunmao-ui/chakra-ui-lib";

type Props = {
  application: any;
  onStoreChange(store: string): void;
  handlers: string[];
  ws: WebSocket;
};

type ServerMessage  = {
  type: string,
  componentId: string;
  name: string;
  parameters?: any;
}

function App(props: Props) {
  const { application, onStoreChange, handlers, ws } = props;
  const { App: SunmaoApp, stateManager, apiService } = initSunmaoUI({
    libs: [
      sunmaoChakraUILib,
      {
        utilMethods: [
          () => {
            return handlers.map((handler) =>
              implementUtilMethod({
                version: "binding/v1",
                metadata: {
                  name: handler,
                },
                spec: {
                  parameters: {} as any,
                },
              })((params) => {
                ws.send(
                    JSON.stringify({
                      type: "Action",
                      handler,
                      params,
                    })
                );
              })
            );
          },
        ],
      },
    ],
  });

  useEffect(() => {
    return watch(stateManager.store, () => {
      onStoreChange(JSON.stringify(stateManager.store));
    });
  }, [stateManager]);

  useEffect(() => {
    const messageHandler = (evt: MessageEvent) => {
      try {
        const message: ServerMessage = JSON.parse((evt.data))
        if (message.type !== 'UiMethod') {
          return
        }
        apiService.send("uiMethod", {
          componentId: message.componentId,
          name: message.name,
          parameters: message.parameters
        })
      } catch (error) {
        console.log("message handler", error)
      }
    }
    ws.addEventListener("message", messageHandler)
    return () => ws.removeEventListener("message", messageHandler)
  }, [apiService])

  return <SunmaoApp options={application} />;
}

export default App;
