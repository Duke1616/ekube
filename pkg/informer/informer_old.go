package informer

//var sharedInformerFactory informers.SharedInformerFactory
//
//func NewSharedInformerFactory(stopCh <-chan struct{}) (err error) {
//	// 1. 加载客户端
//	client, err := resource.NewClient("")
//	if err != nil {
//		return err
//	}
//
//	// 2. 实例化 sharedInformerFactory
//	sharedInformerFactory = informers.NewSharedInformerFactory(client.ClientSet(), time.Second*60)
//
//	// 3. 启动 informer
//	gvrs := []schema.GroupVersionResource{
//		{Group: "", Version: "v1", Resource: "pods"},
//		{Group: "", Version: "v1", Resource: "services"},
//		{Group: "", Version: "v1", Resource: "namespaces"},
//		{Group: "", Version: "v1", Resource: "nodes"},
//
//		{Group: "apps", Version: "v1", Resource: "deployments"},
//		{Group: "apps", Version: "v1", Resource: "statefulsets"},
//		{Group: "apps", Version: "v1", Resource: "daemonsets"},
//	}
//
//	for _, v := range gvrs {
//		// 创建 informer
//		_, err = sharedInformerFactory.ForResource(v)
//		if err != nil {
//			return
//		}
//	}
//
//	// 启动所有创建的 informer
//	sharedInformerFactory.Start(stopCh)
//
//	// 等待所有 informer 全量同步数据完成
//	sharedInformerFactory.WaitForCacheSync(stopCh)
//
//	return
//
//}
//
//func GetInformer() informers.SharedInformerFactory {
//	return sharedInformerFactory
//}
//
//func Setup(stopCh <-chan struct{}) (err error) {
//	err = NewSharedInformerFactory(stopCh)
//	if err != nil {
//		return
//	}
//	return
//}
