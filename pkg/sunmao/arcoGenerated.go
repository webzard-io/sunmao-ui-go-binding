
package sunmao

import (
	"encoding/json"
)


type ArcoTagComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTagComponentBuilder]
}

func (b *ArcoAppBuilder) NewTag() *ArcoTagComponentBuilder {
  t := &ArcoTagComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTagComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"content":"tag","closable":false,"checkable":false,"defaultChecked":false,"color":"","size":"large","bordered":false,"defaultVisible":true}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/tag").Properties(result)
}

type ArcoPaginationComponentBuilder struct {
  *InnerComponentBuilder[*ArcoPaginationComponentBuilder]
}

func (b *ArcoAppBuilder) NewPagination() *ArcoPaginationComponentBuilder {
  t := &ArcoPaginationComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoPaginationComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"pageSize":10,"total":300,"defaultCurrent":3,"disabled":false,"hideOnSinglePage":true,"size":"default","sizeCanChange":false,"simple":false,"showJumper":false,"showTotal":false,"updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/pagination").Properties(result)
}

type ArcoStepsComponentBuilder struct {
  *InnerComponentBuilder[*ArcoStepsComponentBuilder]
}

func (b *ArcoAppBuilder) NewSteps() *ArcoStepsComponentBuilder {
  t := &ArcoStepsComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoStepsComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"type":"default","size":"default","direction":"horizontal","status":"finish","current":2,"lineless":false,"items":[{"title":"Succeeded","description":"This is a description"},{"title":"Processing","description":""},{"title":"Pending","description":"This is a description"}]}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/steps").Properties(result)
}

type ArcoTreeComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTreeComponentBuilder]
}

func (b *ArcoAppBuilder) NewTree() *ArcoTreeComponentBuilder {
  t := &ArcoTreeComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTreeComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"multiple":false,"size":"medium","autoExpandParent":true,"defaultExpandKeys":[],"autoExpandParentWhenDataChanges":false,"data":[{"title":"Asia","key":"asia","children":[{"title":"China","key":"China","selectable":false,"children":[{"title":"Guangdong","key":"Guangdong","children":[{"title":"Guangzhou","key":"Guangzhou"},{"title":"Shenzhen","key":"Shenzhen"}]}]}]},{"title":"Europe","key":"Europe","children":[{"title":"France","key":"France"},{"title":"Germany","key":"Germany"}]}]}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/tree").Properties(result)
}

type ArcoTreeSelectComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTreeSelectComponentBuilder]
}

func (b *ArcoAppBuilder) NewTreeSelect() *ArcoTreeSelectComponentBuilder {
  t := &ArcoTreeSelectComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTreeSelectComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"multiple":false,"defaultValue":["node1"],"treeData":[{"key":"node1","title":"Trunk","disabled":true,"children":[{"key":"node2","title":"Leaf1"}]},{"key":"node3","title":"Trunk2","disabled":false,"children":[{"key":"node4","title":"Leaf2"},{"key":"node5","title":"Leaf3"}]}],"bordered":true,"placeholder":"Select option(s)","labelInValue":true,"size":"default","disabled":false,"error":false,"showSearch":true,"loading":false,"allowClear":true,"updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/treeSelect").Properties(result)
}

type ArcoModalComponentBuilder struct {
  *InnerComponentBuilder[*ArcoModalComponentBuilder]
}

func (b *ArcoAppBuilder) NewModal() *ArcoModalComponentBuilder {
  t := &ArcoModalComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoModalComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"title":"Modal title","mask":true,"simple":false,"okText":"confirm","cancelText":"cancel","closable":true,"maskClosable":true,"confirmLoading":false,"defaultOpen":true,"unmountOnExit":true}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/modal").Properties(result)
}

type ArcoButtonComponentBuilder struct {
  *InnerComponentBuilder[*ArcoButtonComponentBuilder]
}

func (b *ArcoAppBuilder) NewButton() *ArcoButtonComponentBuilder {
  t := &ArcoButtonComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoButtonComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"type":"default","status":"default","long":false,"size":"default","disabled":false,"loading":false,"shape":"square","text":"button"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/button").Properties(result)
}

type ArcoLayoutComponentBuilder struct {
  *InnerComponentBuilder[*ArcoLayoutComponentBuilder]
}

func (b *ArcoAppBuilder) NewLayout() *ArcoLayoutComponentBuilder {
  t := &ArcoLayoutComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoLayoutComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"showHeader":true,"showSideBar":true,"sidebarCollapsible":false,"sidebarDefaultCollapsed":false,"showFooter":true}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/layout").Properties(result)
}

type ArcoImageComponentBuilder struct {
  *InnerComponentBuilder[*ArcoImageComponentBuilder]
}

func (b *ArcoAppBuilder) NewImage() *ArcoImageComponentBuilder {
  t := &ArcoImageComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoImageComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"src":"https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a8c8cdb109cb051163646151a4a5083b.png~tplv-uwbnlip3yd-webp.webp","title":"","description":"","footerPosition":"inner","preview":false,"width":200,"height":200,"error":""}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/image").Properties(result)
}

type ArcoImageGroupComponentBuilder struct {
  *InnerComponentBuilder[*ArcoImageGroupComponentBuilder]
}

func (b *ArcoAppBuilder) NewImageGroup() *ArcoImageGroupComponentBuilder {
  t := &ArcoImageGroupComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoImageGroupComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"imageItems":[{"src":"//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a8c8cdb109cb051163646151a4a5083b.png~tplv-uwbnlip3yd-webp.webp","width":200,"height":200},{"src":"//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/e278888093bef8910e829486fb45dd69.png~tplv-uwbnlip3yd-webp.webp","width":200,"height":200},{"src":"//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp","width":200,"height":200},{"src":"//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/8361eeb82904210b4f55fab888fe8416.png~tplv-uwbnlip3yd-webp.webp","width":200,"height":200}],"src":"//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/8361eeb82904210b4f55fab888fe8416.png~tplv-uwbnlip3yd-webp.webp","infinite":true,"maskClosable":true,"closable":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/imageGroup").Properties(result)
}

type ArcoSelectComponentBuilder struct {
  *InnerComponentBuilder[*ArcoSelectComponentBuilder]
}

func (b *ArcoAppBuilder) NewSelect() *ArcoSelectComponentBuilder {
  t := &ArcoSelectComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoSelectComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"allowClear":false,"multiple":false,"allowCreate":false,"bordered":true,"defaultValue":"Beijing","disabled":false,"labelInValue":false,"loading":false,"showSearch":false,"unmountOnExit":true,"showTitle":false,"options":[{"value":"Beijing","text":"Beijing"},{"value":"London","text":"London"},{"value":"NewYork","text":"NewYork"}],"placeholder":"Select city","size":"default","error":false,"updateWhenDefaultValueChanges":false,"autoFixPosition":false,"autoAlignPopupMinWidth":false,"autoAlignPopupWidth":true,"autoFitPosition":false,"position":"bottom","mountToBody":true}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/select").Properties(result)
}

type ArcoMenuComponentBuilder struct {
  *InnerComponentBuilder[*ArcoMenuComponentBuilder]
}

func (b *ArcoAppBuilder) NewMenu() *ArcoMenuComponentBuilder {
  t := &ArcoMenuComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoMenuComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"mode":"vertical","autoOpen":false,"collapse":false,"autoScrollIntoView":false,"hasCollapseButton":false,"items":[{"key":"key1","text":"item1"},{"key":"key2","text":"item2"}],"ellipsis":false,"defaultActiveKey":"key1","updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/menu").Properties(result)
}

type ArcoDropdownComponentBuilder struct {
  *InnerComponentBuilder[*ArcoDropdownComponentBuilder]
}

func (b *ArcoAppBuilder) NewDropdown() *ArcoDropdownComponentBuilder {
  t := &ArcoDropdownComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoDropdownComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"text":"Click","dropdownType":"default","trigger":"click","position":"bl","disabled":false,"defaultPopupVisible":false,"list":[{"key":"1","label":"Menu item 1"},{"key":"2","label":"Menu item 2"},{"key":"3","label":"Menu item 3"}],"autoAlignPopupWidth":true,"unmountOnExit":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/dropdown").Properties(result)
}

type ArcoSpaceComponentBuilder struct {
  *InnerComponentBuilder[*ArcoSpaceComponentBuilder]
}

func (b *ArcoAppBuilder) NewSpace() *ArcoSpaceComponentBuilder {
  t := &ArcoSpaceComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoSpaceComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"align":"center","direction":"vertical","wrap":false,"size":24}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/space").Properties(result)
}

type ArcoIconComponentBuilder struct {
  *InnerComponentBuilder[*ArcoIconComponentBuilder]
}

func (b *ArcoAppBuilder) NewIcon() *ArcoIconComponentBuilder {
  t := &ArcoIconComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoIconComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"name":"IconArrowUp","spin":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/icon").Properties(result)
}

type ArcoInputComponentBuilder struct {
  *InnerComponentBuilder[*ArcoInputComponentBuilder]
}

func (b *ArcoAppBuilder) NewInput() *ArcoInputComponentBuilder {
  t := &ArcoInputComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoInputComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"allowClear":false,"disabled":false,"readOnly":false,"defaultValue":"","updateWhenDefaultValueChanges":false,"placeholder":"please input","error":false,"size":"default"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/input").Properties(result)
}

type ArcoDividerComponentBuilder struct {
  *InnerComponentBuilder[*ArcoDividerComponentBuilder]
}

func (b *ArcoAppBuilder) NewDivider() *ArcoDividerComponentBuilder {
  t := &ArcoDividerComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoDividerComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"type":"horizontal","orientation":"center"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/divider").Properties(result)
}

type ArcoAvatarComponentBuilder struct {
  *InnerComponentBuilder[*ArcoAvatarComponentBuilder]
}

func (b *ArcoAppBuilder) NewAvatar() *ArcoAvatarComponentBuilder {
  t := &ArcoAvatarComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoAvatarComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"shape":"circle","triggerType":"button","size":50,"type":"text","text":"A","src":"//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/8361eeb82904210b4f55fab888fe8416.png~tplv-uwbnlip3yd-webp.webp"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/avatar").Properties(result)
}

type ArcoMentionsComponentBuilder struct {
  *InnerComponentBuilder[*ArcoMentionsComponentBuilder]
}

func (b *ArcoAppBuilder) NewMentions() *ArcoMentionsComponentBuilder {
  t := &ArcoMentionsComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoMentionsComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"defaultValue":"option1","options":["option1","option2","option3"],"prefix":"@","position":"bl","split":" ","error":false,"allowClear":true,"disabled":false,"placeholder":"you can mentions sb by prefix \"@\"","updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/mentions").Properties(result)
}

type ArcoProgressComponentBuilder struct {
  *InnerComponentBuilder[*ArcoProgressComponentBuilder]
}

func (b *ArcoAppBuilder) NewProgress() *ArcoProgressComponentBuilder {
  t := &ArcoProgressComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoProgressComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"type":"line","status":"normal","color":"#3c92dc","trailColor":"","showText":true,"percent":20,"width":100,"size":"default","animation":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/progress").Properties(result)
}

type ArcoBadgeComponentBuilder struct {
  *InnerComponentBuilder[*ArcoBadgeComponentBuilder]
}

func (b *ArcoAppBuilder) NewBadge() *ArcoBadgeComponentBuilder {
  t := &ArcoBadgeComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoBadgeComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"text":"","dot":false,"count":1,"maxCount":99,"offset":[6,-2]}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/badge").Properties(result)
}

type ArcoTooltipComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTooltipComponentBuilder]
}

func (b *ArcoAppBuilder) NewTooltip() *ArcoTooltipComponentBuilder {
  t := &ArcoTooltipComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTooltipComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"color":"#bbb","position":"bottom","mini":false,"disabled":false,"content":"This is tooltip","trigger":"hover","controlled":false,"unmountOnExit":true}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/tooltip").Properties(result)
}

type ArcoPopoverComponentBuilder struct {
  *InnerComponentBuilder[*ArcoPopoverComponentBuilder]
}

func (b *ArcoAppBuilder) NewPopover() *ArcoPopoverComponentBuilder {
  t := &ArcoPopoverComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoPopoverComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"color":"#eee","position":"bottom","disabled":false,"controlled":false,"trigger":"hover","title":"Title","unmountOnExit":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/popover").Properties(result)
}

type ArcoCollapseComponentBuilder struct {
  *InnerComponentBuilder[*ArcoCollapseComponentBuilder]
}

func (b *ArcoAppBuilder) NewCollapse() *ArcoCollapseComponentBuilder {
  t := &ArcoCollapseComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoCollapseComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"defaultActiveKey":["1"],"options":[{"key":"1","header":"header 1","disabled":false,"showExpandIcon":true},{"key":"2","header":"header 2","disabled":false,"showExpandIcon":true}],"updateWhenDefaultValueChanges":false,"accordion":false,"expandIconPosition":"left","bordered":false,"destroyOnHide":false,"lazyLoad":true}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/collapse").Properties(result)
}

type ArcoCascaderComponentBuilder struct {
  *InnerComponentBuilder[*ArcoCascaderComponentBuilder]
}

func (b *ArcoAppBuilder) NewCascader() *ArcoCascaderComponentBuilder {
  t := &ArcoCascaderComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoCascaderComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"defaultValue":["Zhejiang","Hangzhou","The West Lake"],"expandTrigger":"click","multiple":false,"placeholder":"Please select ...","bordered":true,"size":"default","showSearch":true,"disabled":false,"loading":false,"allowClear":true,"allowCreate":true,"maxTagCount":99,"options":[["Beijing","Dongcheng District","The Forbidden City"],["Shanghai","Pukou","Disney"],["Zhejiang","Hangzhou","The West Lake"]]}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/cascader").Properties(result)
}

type ArcoSkeletonComponentBuilder struct {
  *InnerComponentBuilder[*ArcoSkeletonComponentBuilder]
}

func (b *ArcoAppBuilder) NewSkeleton() *ArcoSkeletonComponentBuilder {
  t := &ArcoSkeletonComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoSkeletonComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"animation":true,"loading":true,"image":false,"imageProps":{"shape":"square","size":"default","position":"left"},"text":true,"textProps":{"rows":3,"width":["100%",600,400]}}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/skeleton").Properties(result)
}

type ArcoTimelineComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTimelineComponentBuilder]
}

func (b *ArcoAppBuilder) NewTimeline() *ArcoTimelineComponentBuilder {
  t := &ArcoTimelineComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTimelineComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"reverse":false,"direction":"vertical","mode":"alternate","labelPosition":"same","items":[{"label":"2017-03-10","content":"The first milestone","dotColor":"#00B42A","lineType":"dashed","lineColor":"#00B42A","dotType":"hollow"},{"label":"2018-05-12","content":"The second milestone","dotColor":"","lineType":"solid","lineColor":"","dotType":"hollow"},{"label":"2020-06-22","content":"The third milestone","dotColor":"#F53F3F","lineType":"dotted","dotType":"solid","lineColor":""},{"label":"2020-09-30","content":"The fourth milestone","dotColor":"#C9CDD4","lineType":"dotted","dotType":"solid","lineColor":""}]}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/timeline").Properties(result)
}

type ArcoRadioComponentBuilder struct {
  *InnerComponentBuilder[*ArcoRadioComponentBuilder]
}

func (b *ArcoAppBuilder) NewRadio() *ArcoRadioComponentBuilder {
  t := &ArcoRadioComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoRadioComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"options":[{"label":"A","value":"a"},{"label":"B","value":"b"},{"label":"C","value":"c"}],"type":"radio","defaultCheckedValue":"a","direction":"horizontal","size":"default","updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/radio").Properties(result)
}

type ArcoCheckboxComponentBuilder struct {
  *InnerComponentBuilder[*ArcoCheckboxComponentBuilder]
}

func (b *ArcoAppBuilder) NewCheckbox() *ArcoCheckboxComponentBuilder {
  t := &ArcoCheckboxComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoCheckboxComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"options":[{"label":"checkbox1","value":"checkbox1"},{"label":"checkbox2","value":"checkbox2"},{"label":"checkbox3","value":"checkbox3"}],"disabled":false,"direction":"horizontal","defaultCheckedValues":["checkbox1"],"showCheckAll":false,"checkAllText":"Check all","updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/checkbox").Properties(result)
}

type ArcoAlertComponentBuilder struct {
  *InnerComponentBuilder[*ArcoAlertComponentBuilder]
}

func (b *ArcoAppBuilder) NewAlert() *ArcoAlertComponentBuilder {
  t := &ArcoAlertComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoAlertComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"closable":true,"title":"info","content":"Here is an example text","showIcon":true,"banner":false,"type":"info"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/alert").Properties(result)
}

type ArcoLinkComponentBuilder struct {
  *InnerComponentBuilder[*ArcoLinkComponentBuilder]
}

func (b *ArcoAppBuilder) NewLink() *ArcoLinkComponentBuilder {
  t := &ArcoLinkComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoLinkComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"asLink":true,"disabled":false,"hoverable":true,"status":"default","href":"#","content":"Link","openInNewPage":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/link").Properties(result)
}

type ArcoSwitchComponentBuilder struct {
  *InnerComponentBuilder[*ArcoSwitchComponentBuilder]
}

func (b *ArcoAppBuilder) NewSwitch() *ArcoSwitchComponentBuilder {
  t := &ArcoSwitchComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoSwitchComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"defaultChecked":false,"disabled":false,"loading":false,"type":"circle","size":"default","updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/switch").Properties(result)
}

type ArcoPasswordInputComponentBuilder struct {
  *InnerComponentBuilder[*ArcoPasswordInputComponentBuilder]
}

func (b *ArcoAppBuilder) NewPasswordInput() *ArcoPasswordInputComponentBuilder {
  t := &ArcoPasswordInputComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoPasswordInputComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"disabled":false,"placeholder":"please input","error":false,"size":"default","visibilityToggle":true,"updateWhenDefaultValueChanges":false,"defaultValue":""}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/passwordInput").Properties(result)
}

type ArcoTextAreaComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTextAreaComponentBuilder]
}

func (b *ArcoAppBuilder) NewTextArea() *ArcoTextAreaComponentBuilder {
  t := &ArcoTextAreaComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTextAreaComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"allowClear":false,"disabled":false,"defaultValue":"","placeholder":"please input","error":false,"size":"default","autoSize":true,"updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/textArea").Properties(result)
}

type ArcoFormControlComponentBuilder struct {
  *InnerComponentBuilder[*ArcoFormControlComponentBuilder]
}

func (b *ArcoAppBuilder) NewFormControl() *ArcoFormControlComponentBuilder {
  t := &ArcoFormControlComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoFormControlComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"label":{"format":"plain","raw":"label"},"layout":"horizontal","required":false,"hidden":false,"extra":"","errorMsg":"","labelAlign":"left","colon":false,"help":"","labelCol":{"span":3,"offset":0},"wrapperCol":{"span":21,"offset":0}}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/formControl").Properties(result)
}

type ArcoDescriptionsComponentBuilder struct {
  *InnerComponentBuilder[*ArcoDescriptionsComponentBuilder]
}

func (b *ArcoAppBuilder) NewDescriptions() *ArcoDescriptionsComponentBuilder {
  t := &ArcoDescriptionsComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoDescriptionsComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"data":[{"label":"Name","value":"Socrates"},{"label":"Mobile","value":"123-1234-1234"},{"label":"Residence","value":"Beijing"},{"label":"Hometown","value":"Beijing"},{"label":"Date of Birth","value":"2020-05-15","span":2},{"label":"Address","value":"Zhichun Road, Beijing"}],"title":"User Info","colon":"","layout":"horizontal","size":"default","tableLayout":"auto","border":true,"column":2}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/descriptions").Properties(result)
}

type ArcoRowComponentBuilder struct {
  *InnerComponentBuilder[*ArcoRowComponentBuilder]
}

func (b *ArcoAppBuilder) NewRow() *ArcoRowComponentBuilder {
  t := &ArcoRowComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoRowComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"gutter":16,"align":"start","justify":"start"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/row").Properties(result)
}

type ArcoColComponentBuilder struct {
  *InnerComponentBuilder[*ArcoColComponentBuilder]
}

func (b *ArcoAppBuilder) NewCol() *ArcoColComponentBuilder {
  t := &ArcoColComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoColComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"offset":0,"pull":0,"push":0,"span":12,"order":0}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/col").Properties(result)
}

type ArcoSliderComponentBuilder struct {
  *InnerComponentBuilder[*ArcoSliderComponentBuilder]
}

func (b *ArcoAppBuilder) NewSlider() *ArcoSliderComponentBuilder {
  t := &ArcoSliderComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoSliderComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"min":0,"max":100,"disabled":false,"tooltipVisible":true,"range":false,"vertical":false,"marks":{},"onlyMarkValue":false,"reverse":false,"step":1,"showTicks":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/slider").Properties(result)
}

type ArcoDatePickerComponentBuilder struct {
  *InnerComponentBuilder[*ArcoDatePickerComponentBuilder]
}

func (b *ArcoAppBuilder) NewDatePicker() *ArcoDatePickerComponentBuilder {
  t := &ArcoDatePickerComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoDatePickerComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"disabled":false,"useUtcOffset":false,"utcOffset":0,"timezone":"","placeholder":"Please Select","position":"bl","dayStartOfWeek":0,"allowClear":false,"type":"date","range":false,"defaultValue":"","showTime":false,"panelOnly":false,"size":"default","disabledTime":{"disabledHours":[0,1,2],"disabledMinutes":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29],"disabledSeconds":[10]},"rangePlaceholder":["Start date","End Date"],"disabledRangeTime":{"disabledHours":[[0,1,2],[7,9,10]],"disabledMinutes":[[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29],[58,59]],"disabledSeconds":[[],[]]},"disabledDate":{"disabledType":"range","dateRange":["2022-5-23","2022-5-26"]},"rangeDisabled":[false,false],"clearRangeOnReselect":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/datePicker").Properties(result)
}

type ArcoTimePickerComponentBuilder struct {
  *InnerComponentBuilder[*ArcoTimePickerComponentBuilder]
}

func (b *ArcoAppBuilder) NewTimePicker() *ArcoTimePickerComponentBuilder {
  t := &ArcoTimePickerComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoTimePickerComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"disabled":false,"disableConfirm":false,"hideDisabledOptions":false,"step":{"hour":1,"minute":1,"second":1},"showNowBtn":true,"placeholder":"Please Select","position":"bl","format":"HH:mm:ss","allowClear":false,"defaultValue":"18:24:23","rangeDefaultValue":["09:24:53","18:44:33"],"range":false,"size":"default","use12Hours":false,"disabledTime":{"disabledHours":[0,1,2],"disabledMinutes":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29],"disabledSeconds":[10]},"rangePlaceholder":["Start date","End Date"],"useUtcOffset":false,"order":false,"utcOffset":0,"timezone":""}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/timePicker").Properties(result)
}

type ArcoNumberInputComponentBuilder struct {
  *InnerComponentBuilder[*ArcoNumberInputComponentBuilder]
}

func (b *ArcoAppBuilder) NewNumberInput() *ArcoNumberInputComponentBuilder {
  t := &ArcoNumberInputComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoNumberInputComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"defaultValue":1,"disabled":false,"placeholder":"please input","error":false,"size":"default","buttonMode":false,"min":0,"max":5,"readOnly":false,"step":1,"precision":1,"updateWhenDefaultValueChanges":false}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/numberInput").Properties(result)
}

type ArcoCarouselComponentBuilder struct {
  *InnerComponentBuilder[*ArcoCarouselComponentBuilder]
}

func (b *ArcoAppBuilder) NewCarousel() *ArcoCarouselComponentBuilder {
  t := &ArcoCarouselComponentBuilder{
    InnerComponentBuilder: newInnerComponent[*ArcoCarouselComponentBuilder](b.AppBuilder),
  }
  t.inner = t
  var jsonBlob = []byte(`{"imageSrc":["//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/cd7a1aaea8e1c5e3d26fe2591e561798.png~tplv-uwbnlip3yd-webp.webp","//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/6480dbc69be1b5de95010289787d64f1.png~tplv-uwbnlip3yd-webp.webp","//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/0265a04fddbd77a19602a15d9d55d797.png~tplv-uwbnlip3yd-webp.webp","//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/24e0dd27418d2291b65db1b21aa62254.png~tplv-uwbnlip3yd-webp.webp"],"indicatorPosition":"bottom","indicatorType":"dot","showArrow":"always","direction":"horizontal","animation":"slide","autoPlay":false,"autoPlayInterval":3000,"autoPlayHoverToPause":true,"moveSpeed":500,"timingFunc":"cubic-bezier(0.34, 0.69, 0.1, 1)"}`)

  result := map[string]interface{}{}
  json.Unmarshal(jsonBlob, &result)

  return t.Type("arco/v1/carousel").Properties(result)
}

