package istio

//
//func control() {
//	for {
//		//获取新创建的Pod
//		pod := client.GetLaatestPod()
//		//Diff以下,检查是否已经初始化过
//		if !isInitialized(pod) {
//			//没有就初始化
//			doInitial(pod)
//		}
//	}
//}
//
//func doInitial(pod interface{}) {
//	cm := client.Get(ConfigMap, "envoy-initializer")
//	newPod := Pod{}
//	newPod.Spec.Containers = cm.Containers
//	newPod.Spec.Volumes = cm.Volumes
//	//生成patch数据
//	patchBytes := strategicpatch.CreateTwoWayMergePatch(pod, newPod)
//	//发起PATCH 请求,修改这个Pod对象
//	client.Patch(pod.Name, patchBytes)
//}
