import { initSunmaoUI } from "@sunmao-ui/runtime";
import "@sunmao-ui/arco-lib/dist/index.css";
import {
  getLibs,
  useApiService,
  BaseProps,
  patchApp,
  patchModules,
} from "./shared";
import { RuntimeModule } from "@sunmao-ui/core";

function App(props: BaseProps) {
  const {
    application,
    modules,
    handlers,
    ws,
    utilMethods,
    applicationPatch,
    modulesPatch,
  } = props;
  const {
    App: SunmaoApp,
    apiService,
    registry,
  } = initSunmaoUI({
    libs: getLibs({ ws, handlers, utilMethods }),
  });

  if (modules) {
    patchModules(modules, modulesPatch).forEach((moduleSchema) => {
      registry.registerModule(moduleSchema as RuntimeModule);
    });
  }

  useApiService({ ws, apiService });

  return <SunmaoApp options={patchApp(application, applicationPatch)} />;
}

export default App;
