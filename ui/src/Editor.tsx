import { initSunmaoUIEditor } from "@sunmao-ui/editor";
import {
  getLibs,
  BaseProps,
  saveApp,
  saveModules,
  patchApp,
  patchModules,
} from "./shared";
import "@sunmao-ui/arco-lib/dist/index.css";
import "@sunmao-ui/editor/dist/index.css";

function Editor(props: BaseProps) {
  const {
    application,
    modules,
    handlers,
    ws,
    utilMethods,
    applicationPatch,
    modulesPatch,
  } = props;
  const { Editor } = initSunmaoUIEditor({
    defaultApplication: patchApp(application, applicationPatch),
    defaultModules: patchModules(modules, modulesPatch),
    runtimeProps: {
      libs: getLibs({ ws, handlers, utilMethods }),
    },
    storageHandler: {
      onSaveApp: function (newApp) {
        saveApp(newApp, application);
      },
      onSaveModules: function (newModules) {
        saveModules(newModules, modules || []);
      },
    },
  });

  // TODO: call the useApiService hook when sunmao-ui expose apiService in editor mode

  return <Editor />;
}

export default Editor;
