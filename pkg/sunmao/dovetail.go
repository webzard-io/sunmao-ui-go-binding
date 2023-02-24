package sunmao

import (
	"encoding/json"
)

type DovetailAppBuilder struct {
	*AppBuilder
}

func NewDovetailApp() *DovetailAppBuilder {
	b := &DovetailAppBuilder{
		AppBuilder: NewApp(),
	}
	return b
}

type DovetailRootComponentBuilder struct {
	*InnerComponentBuilder[*DovetailRootComponentBuilder]
}

func (b *DovetailAppBuilder) NewRoot() *DovetailRootComponentBuilder {
	t := &DovetailRootComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailRootComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/root").Properties(result)
}

type DovetailButtonComponentBuilder struct {
	*InnerComponentBuilder[*DovetailButtonComponentBuilder]
}

func (b *DovetailAppBuilder) NewButton() *DovetailButtonComponentBuilder {
	t := &DovetailButtonComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailButtonComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"type":"primary","text":{"raw":"text","format":"plain"}}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/button").Properties(result)
}

type DovetailKubectlGetTableComponentBuilder struct {
	*InnerComponentBuilder[*DovetailKubectlGetTableComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlGetTable() *DovetailKubectlGetTableComponentBuilder {
	t := &DovetailKubectlGetTableComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailKubectlGetTableComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"basePath":"proxy-k8s","apiBase":"apis/kubesmart.smtx.io/v1alpha1","resource":"kubesmartclusters","columns":[{"key":"name","title":"Name","dataIndex":"metadata.name","width":100,"sortType":"auto","filters":[]},{"key":"namespace","title":"Namespace","dataIndex":"metadata.namespace","sortType":"none","width":200,"filters":[]},{"key":"age","title":"Age","dataIndex":"metadata.creationTimestamp","sortType":"auto","width":100,"filters":[]}],"canActive":true,"defaultSize":10,"empty":"No Data."}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/kubectl_get_table").Properties(result)
}

type DovetailObjectAgeComponentBuilder struct {
	*InnerComponentBuilder[*DovetailObjectAgeComponentBuilder]
}

func (b *DovetailAppBuilder) NewObjectAge() *DovetailObjectAgeComponentBuilder {
	t := &DovetailObjectAgeComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailObjectAgeComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"date":"2022-01-01"}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/object_age").Properties(result)
}

type DovetailUnstructuredSidebarComponentBuilder struct {
	*InnerComponentBuilder[*DovetailUnstructuredSidebarComponentBuilder]
}

func (b *DovetailAppBuilder) NewUnstructuredSidebar() *DovetailUnstructuredSidebarComponentBuilder {
	t := &DovetailUnstructuredSidebarComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailUnstructuredSidebarComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"item":{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"name":"nginx-deployment","labels":{"app":"nginx"}},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx:1.14.2","ports":[{"containerPort":80}]}]}}}},"defaultVisible":true}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/unstructured_sidebar").Properties(result)
}

type DovetailUnstructuredPageComponentBuilder struct {
	*InnerComponentBuilder[*DovetailUnstructuredPageComponentBuilder]
}

func (b *DovetailAppBuilder) NewUnstructuredPage() *DovetailUnstructuredPageComponentBuilder {
	t := &DovetailUnstructuredPageComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailUnstructuredPageComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"kind":"Deployment","apiBase":"apis/apps/v1/deployments","defaultVisible":true}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/unstructured_page").Properties(result)
}

type DovetailModalComponentBuilder struct {
	*InnerComponentBuilder[*DovetailModalComponentBuilder]
}

func (b *DovetailAppBuilder) NewModal() *DovetailModalComponentBuilder {
	t := &DovetailModalComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailModalComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"title":"Header"}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/modal").Properties(result)
}

type DovetailSelectComponentBuilder struct {
	*InnerComponentBuilder[*DovetailSelectComponentBuilder]
}

func (b *DovetailAppBuilder) NewSelect() *DovetailSelectComponentBuilder {
	t := &DovetailSelectComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailSelectComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"options":[{"text":"option 1","value":1},{"text":"option 2","value":"2"}]}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/select").Properties(result)
}

type DovetailIconComponentBuilder struct {
	*InnerComponentBuilder[*DovetailIconComponentBuilder]
}

func (b *DovetailAppBuilder) NewIcon() *DovetailIconComponentBuilder {
	t := &DovetailIconComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailIconComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"name":"ExpandAltOutlined"}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/icon").Properties(result)
}

type DovetailCodeEditorComponentBuilder struct {
	*InnerComponentBuilder[*DovetailCodeEditorComponentBuilder]
}

func (b *DovetailAppBuilder) NewCodeEditor() *DovetailCodeEditorComponentBuilder {
	t := &DovetailCodeEditorComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailCodeEditorComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"defaultValue":"console.log('hello world')","language":"javascript"}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/code_editor").Properties(result)
}

type DovetailKubectlApplyFormComponentBuilder struct {
	*InnerComponentBuilder[*DovetailKubectlApplyFormComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlApplyForm() *DovetailKubectlApplyFormComponentBuilder {
	t := &DovetailKubectlApplyFormComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailKubectlApplyFormComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"applyConfig":{"create":true,"patch":true},"formConfig":{"yaml":"","schemas":[],"defaultValues":[],"uiConfig":{"layout":{"type":"simple","fields":[]}}}}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/kubectl_apply_form").Properties(result)
}

type DovetailKubectlGetDetailComponentBuilder struct {
	*InnerComponentBuilder[*DovetailKubectlGetDetailComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlGetDetail() *DovetailKubectlGetDetailComponentBuilder {
	t := &DovetailKubectlGetDetailComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailKubectlGetDetailComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"basePath":"proxy-k8s","apiBase":"/apis/apps/v1","namespace":"kube-system","resource":"deployments","name":"coredns","layout":{"type":"tabs","tabs":[{"label":"Detail","key":"detail","sections":[{"title":"","info":{"metadata":[{"label":"Name","path":"metadata.name"},{"label":"Labels","path":"metadata.labels"},{"label":"Age","path":"metadata.creationTimestamp"}]}}]},{"label":"Monitor","key":"monitor","sections":[]}],"sections":[]}}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/kubectl_get_detail").Properties(result)
}

type DovetailKubectlGetListComponentBuilder struct {
	*InnerComponentBuilder[*DovetailKubectlGetListComponentBuilder]
}

func (b *DovetailAppBuilder) NewKubectlGetList() *DovetailKubectlGetListComponentBuilder {
	t := &DovetailKubectlGetListComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailKubectlGetListComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"basePath":"proxy-k8s","apiBase":"apis/kubesmart.smtx.io/v1alpha1","resource":"kubesmartclusters"}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/kubectl_get_list").Properties(result)
}

type DovetailModalV2ComponentBuilder struct {
	*InnerComponentBuilder[*DovetailModalV2ComponentBuilder]
}

func (b *DovetailAppBuilder) NewV2Modal() *DovetailModalV2ComponentBuilder {
	t := &DovetailModalV2ComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailModalV2ComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"title":"Header","zIndex":1000}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v2/modal").Properties(result)
}

type DovetailConfirmModalComponentBuilder struct {
	*InnerComponentBuilder[*DovetailConfirmModalComponentBuilder]
}

func (b *DovetailAppBuilder) NewConfirmModal() *DovetailConfirmModalComponentBuilder {
	t := &DovetailConfirmModalComponentBuilder{
		InnerComponentBuilder: newInnerComponent[*DovetailConfirmModalComponentBuilder](b.AppBuilder),
	}
	t.inner = t
	var jsonBlob = []byte(`{"width":492,"title":"Delete","errors":[],"size":"small"}`)

	result := map[string]interface{}{}
	json.Unmarshal(jsonBlob, &result)

	return t.Type("kui/v1/confirm_modal").Properties(result)
}
