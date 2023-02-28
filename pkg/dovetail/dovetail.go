
  package dovetail
  
  import (
    "encoding/json"
    "github.com/yuyz0112/sunmao-ui-go-binding/pkg/sunmao"
  )
  
  type DovetailAppBuilder struct {
    *sunmao.AppBuilder
  }
  
  func NewDovetailApp(appBuilder *sunmao.AppBuilder) *DovetailAppBuilder {
    b := &DovetailAppBuilder{
      appBuilder,
    }
    return b
  }
  
type DovetailRootComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailRootComponentBuilder]
}

func (b *DovetailAppBuilder) NewRoot() *DovetailRootComponentBuilder {
	c := &DovetailRootComponentBuilder{
		sunmao.NewInnerComponent[*DovetailRootComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/root").Properties(result)
}

type DovetailButtonComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailButtonComponentBuilder]
}

func (b *DovetailAppBuilder) NewButton() *DovetailButtonComponentBuilder {
	c := &DovetailButtonComponentBuilder{
		sunmao.NewInnerComponent[*DovetailButtonComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"type":"primary","text":{"raw":"text","format":"plain"}}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/button").Properties(result)
}

type DovetailKubectlGetTableComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailKubectlGetTableComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlGetTable() *DovetailKubectlGetTableComponentBuilder {
	c := &DovetailKubectlGetTableComponentBuilder{
		sunmao.NewInnerComponent[*DovetailKubectlGetTableComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"basePath":"proxy-k8s","apiBase":"apis/kubesmart.smtx.io/v1alpha1","resource":"kubesmartclusters","columns":[{"key":"name","title":"Name","dataIndex":"metadata.name","width":100,"sortType":"auto","filters":[]},{"key":"namespace","title":"Namespace","dataIndex":"metadata.namespace","sortType":"none","width":200,"filters":[]},{"key":"age","title":"Age","dataIndex":"metadata.creationTimestamp","sortType":"auto","width":100,"filters":[]}],"canActive":true,"defaultSize":10,"empty":"No Data."}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/kubectl_get_table").Properties(result)
}

type DovetailObjectAgeComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailObjectAgeComponentBuilder]
}

func (b *DovetailAppBuilder) NewObjectAge() *DovetailObjectAgeComponentBuilder {
	c := &DovetailObjectAgeComponentBuilder{
		sunmao.NewInnerComponent[*DovetailObjectAgeComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"date":"2022-01-01"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/object_age").Properties(result)
}

type DovetailUnstructuredSidebarComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailUnstructuredSidebarComponentBuilder]
}

func (b *DovetailAppBuilder) NewUnstructuredSidebar() *DovetailUnstructuredSidebarComponentBuilder {
	c := &DovetailUnstructuredSidebarComponentBuilder{
		sunmao.NewInnerComponent[*DovetailUnstructuredSidebarComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"item":{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"name":"nginx-deployment","labels":{"app":"nginx"}},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx:1.14.2","ports":[{"containerPort":80}]}]}}}},"defaultVisible":true}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/unstructured_sidebar").Properties(result)
}

type DovetailUnstructuredPageComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailUnstructuredPageComponentBuilder]
}

func (b *DovetailAppBuilder) NewUnstructuredPage() *DovetailUnstructuredPageComponentBuilder {
	c := &DovetailUnstructuredPageComponentBuilder{
		sunmao.NewInnerComponent[*DovetailUnstructuredPageComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"kind":"Deployment","apiBase":"apis/apps/v1/deployments","defaultVisible":true}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/unstructured_page").Properties(result)
}

type DovetailModalComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailModalComponentBuilder]
}

func (b *DovetailAppBuilder) NewModal() *DovetailModalComponentBuilder {
	c := &DovetailModalComponentBuilder{
		sunmao.NewInnerComponent[*DovetailModalComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"title":"Header"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/modal").Properties(result)
}

type DovetailSelectComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailSelectComponentBuilder]
}

func (b *DovetailAppBuilder) NewSelect() *DovetailSelectComponentBuilder {
	c := &DovetailSelectComponentBuilder{
		sunmao.NewInnerComponent[*DovetailSelectComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"options":[{"text":"option 1","value":1},{"text":"option 2","value":"2"}]}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/select").Properties(result)
}

type DovetailIconComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailIconComponentBuilder]
}

func (b *DovetailAppBuilder) NewIcon() *DovetailIconComponentBuilder {
	c := &DovetailIconComponentBuilder{
		sunmao.NewInnerComponent[*DovetailIconComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"name":"ExpandAltOutlined"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/icon").Properties(result)
}

type DovetailCodeEditorComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailCodeEditorComponentBuilder]
}

func (b *DovetailAppBuilder) NewCodeEditor() *DovetailCodeEditorComponentBuilder {
	c := &DovetailCodeEditorComponentBuilder{
		sunmao.NewInnerComponent[*DovetailCodeEditorComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"defaultValue":"console.log('hello world')","language":"javascript"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/code_editor").Properties(result)
}

type DovetailKubectlApplyFormComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailKubectlApplyFormComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlApplyForm() *DovetailKubectlApplyFormComponentBuilder {
	c := &DovetailKubectlApplyFormComponentBuilder{
		sunmao.NewInnerComponent[*DovetailKubectlApplyFormComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"formConfig":{"yaml":"","schemas":[],"defaultValues":[],"uiConfig":{"layout":{"type":"simple","fields":[]}}}}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/kubectl_apply_form").Properties(result)
}

type DovetailKubectlGetDetailComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailKubectlGetDetailComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlGetDetail() *DovetailKubectlGetDetailComponentBuilder {
	c := &DovetailKubectlGetDetailComponentBuilder{
		sunmao.NewInnerComponent[*DovetailKubectlGetDetailComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"basePath":"/api/k8s","apiBase":"/apis/storage.k8s.io/v1","namespace":"","resource":"storageclasses","name":"","layout":{"type":"tabs","tabs":[{"label":"Detail","key":"detail","sections":[{"title":"","info":{"metadata":[{"label":"Name","path":"metadata.name","condition":true},{"label":"Labels","path":"metadata.labels","condition":true},{"label":"Age","path":"metadata.creationTimestamp","condition":true}]}}]},{"label":"Monitor","key":"monitor","sections":[]}],"sections":[]},"query":{}}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/kubectl_get_detail").Properties(result)
}

type DovetailKubectlGetListComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailKubectlGetListComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlGetList() *DovetailKubectlGetListComponentBuilder {
	c := &DovetailKubectlGetListComponentBuilder{
		sunmao.NewInnerComponent[*DovetailKubectlGetListComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"basePath":"proxy-k8s","apiBase":"apis/kubesmart.smtx.io/v1alpha1","resource":"kubesmartclusters"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/kubectl_get_list").Properties(result)
}

type DovetailModalV2ComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailModalV2ComponentBuilder]
}

func (b *DovetailAppBuilder) NewV2Modal() *DovetailModalV2ComponentBuilder {
	c := &DovetailModalV2ComponentBuilder{
		sunmao.NewInnerComponent[*DovetailModalV2ComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"title":"Header","zIndex":1000}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v2/modal").Properties(result)
}

type DovetailConfirmModalComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailConfirmModalComponentBuilder]
}

func (b *DovetailAppBuilder) NewConfirmModal() *DovetailConfirmModalComponentBuilder {
	c := &DovetailConfirmModalComponentBuilder{
		sunmao.NewInnerComponent[*DovetailConfirmModalComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"width":492,"title":"Delete","errors":[],"size":"small"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/confirm_modal").Properties(result)
}

type DovetailK8sLabelGroupComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailK8sLabelGroupComponentBuilder]
}

func (b *DovetailAppBuilder) NewK8sLabelGroup() *DovetailK8sLabelGroupComponentBuilder {
	c := &DovetailK8sLabelGroupComponentBuilder{
		sunmao.NewInnerComponent[*DovetailK8sLabelGroupComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"labels":{},"editable":false}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/k8s_label_group").Properties(result)
}

type DovetailK8sNamespaceSelectComponentBuilder struct {
  *sunmao.InnerComponentBuilder[*DovetailK8sNamespaceSelectComponentBuilder]
}

func (b *DovetailAppBuilder) NewK8sNamespaceSelect() *DovetailK8sNamespaceSelectComponentBuilder {
	c := &DovetailK8sNamespaceSelectComponentBuilder{
		sunmao.NewInnerComponent[*DovetailK8sNamespaceSelectComponentBuilder](b.AppBuilder),
	}
	c.Inner = c

  var jsonBlob = []byte(`{"placeholder":"Please select"}`)
  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return c.Type("kui/v1/k8s_namespace_select").Properties(result)
}

  