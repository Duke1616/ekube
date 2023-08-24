package user

import (
	"context"
	usersv1 "ekube/api/k8s/users/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// UsersReconciler reconciles a Book object
type UsersReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.github.com,resources=books,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.github.com,resources=books/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.github.com,resources=books/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Book object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *UsersReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	instance := &usersv1.Users{}

	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UsersReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.Client == nil {
		r.Client = mgr.GetClient()
	}

	if r.Scheme == nil {
		r.Scheme = mgr.GetScheme()
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&usersv1.Users{}).
		Complete(r)
}
