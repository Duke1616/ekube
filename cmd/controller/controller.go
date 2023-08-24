package controller

import (
	"ekube/cmd/controller/option"
	"ekube/pkg/controller/user"
	"ekube/pkg/informer"
	"ekube/pkg/k8s/client"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var allControllers = []string{
	"user",
}

// setup all available controllers one by one
func addAllControllers(mgr manager.Manager, client client.Client, informerFactory informer.InformerFactory, cmOptions *option.KubeSphereControllerManagerOptions,
	stopCh <-chan struct{}) error {

	if cmOptions.IsControllerEnabled("user") {
		userController := &user.UsersReconciler{}
		addControllerWithSetup(mgr, "users", userController)
	}
	// log all controllers process result
	for _, name := range allControllers {
		if addSuccessfullyControllers.Has(name) {
			klog.Infof("%s controller is enabled and added successfully.", name)
		} else {
			klog.Infof("%s controller is enabled but is not going to run due to its dependent component being disabled.", name)
		}
	}

	return nil
}

type setupAbleController interface {
	SetupWithManager(mgr ctrl.Manager) error
}

func addControllerWithSetup(mgr manager.Manager, name string, controller setupAbleController) {
	if err := controller.SetupWithManager(mgr); err != nil {
		klog.Fatalf("Unable to create %v controller: %v", name, err)
	}
	addSuccessfullyControllers.Insert(name)
}

var addSuccessfullyControllers = sets.New[string]()
