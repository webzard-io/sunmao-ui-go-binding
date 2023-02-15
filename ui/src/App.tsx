import { initSunmaoUI } from "@sunmao-ui/runtime";
import "@sunmao-ui/arco-lib/dist/index.css";
import {
  getLibs,
  useApiService,
  BaseProps,
  patchApp,
  patchModules,
  mergeWithBaseApplication,
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
    applicationBase,
  } = props;
  const {
    App: SunmaoApp,
    apiService,
    registry,
  } = initSunmaoUI({
    libs: getLibs({ ws, handlers, utilMethods }),
  });

  const patchedApp = patchApp(applicationBase || application, applicationPatch);
  let mergedApp = patchedApp;
  if (applicationBase) {
    mergedApp = mergeWithBaseApplication(
      applicationBase,
      application,
      patchedApp
    );
  }

  if (modules) {
    patchModules(modules, modulesPatch).forEach((moduleSchema) => {
      registry.registerModule(moduleSchema as RuntimeModule);
    });
  }

  useApiService({ ws, apiService });

  return <SunmaoApp options={mergedApp} />;
}

export default App;
